package structs

type User struct {
	Id					int			`json:"-" bson:"-"`
	Username			string		`json:"username"`
	Email				string		`json:"email"`
	Password			string		`json:"password,omitempty"`
	Chats				[]Chat		`json:"chats"`
	Pool				string		`json:"poolId" bson:"-"`
	SecretHash			string		`json:"secret_hash" bson:"-"`
}

type Chat struct {
	ChatId				string		`json:"chat_id" bson:"chatid"`
	Name				string		`json:"name" bson:"name"`
	Usernames			[]string	`json:"usernames" bson:"usernames"`
	Admin				string		`json:"admin" bson:"admin"`
	MessagePoolId		string		`json:"-" bson:"messagepoolid"`
	LastReadMessageId	string		`json:"lastReadMessageId" bson:"last_read_message_id"`
}

type Message struct {
	Id					string		`json:"id" bson:"id"`
	Sender				string		`json:"sender" bson:"sender"`
	ChatId				string		`json:"chat_id" bson:"-"`
	Date				int			`json:"date" bson:"date"`
	State				int			`json:"state" bson:"state"`
	Text				string		`json:"text" bson:"text"`
	Meta				int			`json:"meta" bson:"-"`
	AttachedFilePath	string		`json:"attached_file_path" bson:"attachedfilepath"`
}

type Pool struct {
	PoolId				string		`json:"poolId"`
	Password			string		`json:"password"`
}
