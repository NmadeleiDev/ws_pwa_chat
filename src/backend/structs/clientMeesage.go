package structs

type ClientMessage struct {
	Type			int			`json:"type"`
	Message			Message		`json:"message"`
}
