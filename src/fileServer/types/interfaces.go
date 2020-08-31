package types

import "os"

type S3Manager interface {
	UploadFile(filepath string) (string, bool)
	DownloadFile(fileId string) (*os.File, error)
}

type FileInfoStorage interface {
	CloseConnection()

	TrySaveFileIdToLot(fileInfo ClientFileInfo) bool
	GetFileInfoFromLot(lotId, viewToken string) *ClientFileInfo
}