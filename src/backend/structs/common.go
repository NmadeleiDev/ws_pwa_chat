package structs

type User struct {
	Id       string `json:"-" bson:"-"`
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password,omitempty" bson:"-"`
	Chats    []Chat `json:"chats" bson:"chats"`
	IsOnline bool   `json:"isOnline" bson:"is_online"`
	Pool     string `json:"poolId" bson:"pool_id"`
	Token    string `json:"token,omitempty" bson:"-"`
}

type Chat struct {
	ChatId            string   `json:"id" bson:"chatid"`
	Name              string   `json:"name" bson:"name"`
	Usernames         []string `json:"usernames" bson:"usernames"`
	Admin             string   `json:"admin" bson:"admin"`
	MessagePoolId     string   `json:"-" bson:"messagepoolid"`
	LastReadMessageId string   `json:"lastReadMessageId" bson:"last_read_message_id"`
	StorePeriod       int      `json:"storePeriod" bson:"store_period"`
}

type Message struct {
	Id               string `json:"id" bson:"id"`
	Sender           string `json:"sender" bson:"sender"`
	ChatId           string `json:"chatId" bson:"-"`
	Date             int    `json:"date" bson:"date"`
	State            int    `json:"state" bson:"state"`
	Text             string `json:"text" bson:"text"`
	Meta             int    `json:"meta" bson:"-"`
	AttachedFilePath string `json:"attachedFilePath" bson:"attachedfilepath"`
}

type Pool struct {
	PoolId   string `json:"poolId"`
	Password string `json:"password"`
}

type ChatWithMessages struct {
	ChatId            string    `json:"id" bson:"chatid"`
	Name              string    `json:"name" bson:"name"`
	Usernames         []string  `json:"usernames" bson:"usernames"`
	Admin             string    `json:"admin" bson:"admin"`
	MessagePoolId     string    `json:"-" bson:"messagepoolid"`
	LastReadMessageId string    `json:"lastReadMessageId" bson:"last_read_message_id"`
	StorePeriod       int       `json:"storePeriod" bson:"storePeriod"`
	Messages          []Message `json:"messages" bson:"-"`
}

type ChatInfo struct {
	ChatId        string `json:"id" bson:"chatid"`
	MessagePoolId string `json:"-" bson:"messagepoolid"`
}
