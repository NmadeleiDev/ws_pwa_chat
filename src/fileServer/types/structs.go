package types

type HttpResponse struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

type ClientFileInfo struct {
	Id		string
	ContentType	string
	Size		int64
	LotId		int64
}
