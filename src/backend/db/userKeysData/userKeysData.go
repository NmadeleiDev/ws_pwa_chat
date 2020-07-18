package userKeysData

import (
	"chat_backend/db/userKeysData/postgres"
	"chat_backend/structs"
)

var Manager UserKeysData

type UserKeysData interface {
	MakeConnection()
	CloseConnection()
	InitTables()

	SignInAndRefreshMobileAndWebKeys(user structs.User) (webKey string, mobileToken string, ok bool)

	// mobile auth
	IdentifyUserByToken(auth structs.MobileToken, time string) (id string, ok bool)

	// web auth
	IdentifyUserByCookie(cookie string) (id string, ok bool)
	DeleteCookieKey(sessionKey string)

	CreateUserAndGenerateKeys(user structs.User) (webKey string, mobileToken string, ok bool)
	DeleteUser(user structs.User)

	ToggleUserOnlineState(id string, state bool) bool
	GetAllSamePoolUsers(id string) ([]structs.User, error)
	UpdateUserPoolId(userId string, pool structs.Pool) bool
	CreatePool(pool structs.Pool) bool
	TryPoolSignIn(pool structs.Pool) bool

	SetUserSecret(user structs.User) bool
	SetSessionKeyById(sessionKey string, id int) bool
	UpdateSessionKey(old, new string) bool

	GetUserIdByName(name string) (string, bool)
}

func Init() {
	Manager = &postgres.PgSqlUserKeysManager{}
	Manager.MakeConnection()
	Manager.InitTables()
}
