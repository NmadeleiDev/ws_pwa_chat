package fileInfoManager

import (
	_ "crypto/md5"
	"database/sql"
	"fileServer/types"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"os"
)

const (
	filesInfoTable  = "file_server.files_info"
	hashCost       = 14
)

var Manager types.FileInfoStorage

type FileInfoManager struct {
	connection *sql.DB
}

func (db *FileInfoManager) TrySaveFileIdToLot(fileInfo types.ClientFileInfo) bool {
	query := `
UPDATE ` + filesInfoTable + `
SET file_id=$1, lot_status=1, content_type=$2, file_size=$3
WHERE lot_id=$4 AND lot_status=0`
// затираем lot_token мусором сразу после сохранения файла. Таким образом, после этой операции, к нему уже ни у кого не будет доступа, пока главный сервер не поставит туда какой - нибудь новый токен, и не сообщит его авторизованному пользователю

	logrus.Infof("Updating lot: %v", fileInfo)

	if _, err := db.connection.Exec(query,
		fileInfo.Id,
		fileInfo.ContentType,
		fileInfo.Size,
		fileInfo.LotId); err != nil {
		logrus.Errorf("Error saving file info: %v", err)
		return false
	}
	return true
}

func (db *FileInfoManager) GetFileInfoFromLot(lotId, viewToken string) *types.ClientFileInfo {
	fileInfo := types.ClientFileInfo{}

	query := `
SELECT file_id, content_type, file_size
FROM ` + filesInfoTable + ` 
WHERE lot_id=$1 AND $2=ANY(view_tokens)`

	if err := db.connection.QueryRow(query, lotId, viewToken).Scan(&fileInfo.Id, &fileInfo.ContentType, &fileInfo.Size); err != nil {
		logrus.Errorf("Error selecting file id: %v", err)
		return nil
	}

	query = `
UPDATE ` + filesInfoTable + ` 
SET view_tokens=array_remove(view_tokens, $2) 
WHERE lot_id=$1`

	if _, err := db.connection.Exec(query, lotId, viewToken); err != nil {
		logrus.Errorf("Error updating view tokens: %v", err)
		return nil
	}

	return &fileInfo
}

func Init() {
	db := FileInfoManager{}
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	dbName := os.Getenv("POSTGRES_DB")

	connStr := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", user, password, host, port, dbName)
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		logrus.Fatal("Error connecting to database: ", err)
	}
	db.connection = conn
	Manager = &db
}

func (db *FileInfoManager) CloseConnection() {
	if err := db.connection.Close(); err != nil {
		logrus.Error("Error closing fileInfoManager connection: ", err)
	}
}
