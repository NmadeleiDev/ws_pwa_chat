package mongodb

import "strconv"

func getChatsObjName(len int) string {
	chatsObjName := "chats"
	if len > 0 {
		chatsObjName += "." + strconv.Itoa(len)
	}
	return chatsObjName
}
