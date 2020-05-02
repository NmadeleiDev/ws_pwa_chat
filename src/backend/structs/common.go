package structs

type User struct {
	Username			string		`json:"username"`
	Email				string		`json:"email"`
	Password			string		`json:"password,omitempty"`
	Chats				[]Chat		`json:"chats"`
}

type Chat struct {
	ChatId				string		`json:"chat_id" bson:"chatid"`
	Name				string		`json:"name" bson:"name"`
	Usernames			[]string	`json:"usernames" bson:"usernames"`
	Admin				string		`json:"admin" bson:"admin"`
	MessagePoolId		string		`json:"-" bson:"messagepoolid"`
}

type Message struct {
	Sender				string		`json:"sender" bson:"sender"`
	ChatId				string		`json:"chat_id" bson:"chatid"`
	IsRead				bool		`json:"is_read" bson:"isread"`
	Date				int			`json:"date" bson:"date"`
	State				int			`json:"state" bson:"state"`
	Text				string		`json:"text" bson:"text"`
	Meta				int			`json:"meta" bson:"-"`
	AttachedFilePath	string		`json:"attached_file_path" bson:"attachedfilepath"`
}
