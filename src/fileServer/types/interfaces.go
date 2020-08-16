package types

type S3Manager interface {
	Init()

	UploadFile(filepath string) bool
	DownloadFile(fileId string) bool
}

type FileInfoStorage interface {
	MakeConnection()
	CloseConnection()

	TrySaveFileIdToLot(fileId, lotId, lotToken string) bool
	GetFileIdFromLot(lotId, viewToken string) string
}