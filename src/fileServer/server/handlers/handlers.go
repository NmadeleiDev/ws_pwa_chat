package handlers

import (
	"bufio"
	"bytes"
	"fileServer/db/fileInfoManager"
	"fileServer/s3Manager"
	"fileServer/types"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

const (
	maxFileSize = 10 << 22
	formSaveKeyFieldName = "save_key"
	formFileFieldName = "file"
	filesFolder = "./files"
)

func ManagerFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.Body = http.MaxBytesReader(w, r.Body, maxFileSize+1024)
		reader, err := r.MultipartReader()
		if err != nil {
			logrus.Warnf("error creating reader: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		p, err := reader.NextPart()
		if err != nil {
			logrus.Warnf("error reading next part (first): %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if p.FormName() != formSaveKeyFieldName {
			logrus.Warnf("%v is expected, got: %v", formSaveKeyFieldName, p.FormName())
			return
		}
		saveKey := new(bytes.Buffer)
		_, err = saveKey.ReadFrom(p)
		if err != nil && err != io.EOF {
			logrus.Warnf("eol: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		p, err = reader.NextPart()
		if err != nil {
			if err == io.EOF {
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logrus.Warnf("error reading next part: %v", err)
			return
		}
		if p.FormName() != formFileFieldName {
			logrus.Warnf("%v is expected, got: %v", formFileFieldName, p.FormName())
			http.Error(w, "file_field is expected", http.StatusBadRequest)
			return
		}
		buf := bufio.NewReader(p)
		sniff, err := buf.Peek(512)
		if err != nil {
			SendFailResponse(w, "unknown content type")
			return
		}
		fileInfo := types.ClientFileInfo{}
		fileInfo.LotId = parseSaveKeyToInt(saveKey.Bytes())
		fileInfo.ContentType = http.DetectContentType(sniff)
		f, err := ioutil.TempFile(filesFolder, "")
		if err != nil {
			logrus.Warnf("error creating temp file: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer removeFile(f)
		fileInfo.Id = f.Name()

		lmt := io.MultiReader(buf, io.LimitReader(p, maxFileSize - 511))
		fileInfo.Size, err = io.Copy(f, lmt)
		if err != nil && err != io.EOF {
			logrus.Warnf("error reading file: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if fileInfo.Size > maxFileSize {
			os.Remove(fileInfo.Id)
			logrus.Warnf("file size over limit, written: %v", fileInfo.Size)
			http.Error(w, "file size over limit", http.StatusBadRequest)
			return
		}
		logrus.Infof("downloaded, size = %v", fileInfo.Size)

		var ok bool
		fileInfo.Id, ok = s3Manager.Manager.UploadFile(fileInfo.Id)
		if !ok {
			SendFailResponse(w, "file upload error")
			return
		}
		if fileInfoManager.Manager.TrySaveFileIdToLot(fileInfo) {
			SendDataResponse(w, "ok")
		} else {
			SendFailResponse(w, "Error saving file data")
		}
	} else if r.Method == http.MethodGet {
		lotId := r.URL.Query().Get("file")
		key := r.URL.Query().Get("key")
		if len(lotId) == 0 || len(key) == 0 {
			logrus.Errorf("Lot id or key is not set: file = %v; key = %v", lotId, key)
			SendFailResponse(w, "File or key is not set")
			return
		}
		fileInfo := fileInfoManager.Manager.GetFileInfoFromLot(lotId, key)
		if len(fileInfo.Id) == 0 {
			logrus.Errorf("Failed to get file info from lot: lotId = %v; key = %v", lotId, key)
			SendFailResponse(w, "key is invalid")
			return
		}
		file, err := s3Manager.Manager.DownloadFile(fileInfo.Id)
		if err != nil {
			logrus.Errorf("Error getting file: %v", err)
			SendFailResponse(w, "file not found")
			return
		}
		defer func() {
			removeFile(file)
		}()

		b := make([]byte, fileInfo.Size)
		read, err := file.Read(b)
		if err != nil {
			logrus.Errorf("Error reading file: %v ", err)
			SendFailResponse(w, "Error reading file")
			return
		}
		logrus.Infof("Expected %v, got %v", fileInfo.Size, read)

		w.Header().Set("Content-Type", fileInfo.ContentType)
		w.Header().Set("Content-Length", strconv.Itoa(len(b)))
		if _, err := w.Write(b); err != nil {
			logrus.Errorf("unable to write file: %v", err)
			SendFailResponse(w, "unable to write file")
		}
	}
}

