package structs

type HttpResponse struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

type SocketMessage struct {
	Type  int         `json:"type"`
	Data  interface{} `json:"data"`
	Error error       `json:"error"`
}
