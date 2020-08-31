package postgres

import (
	"chat_backend/hashes"
	"chat_backend/structs"
	"crypto/md5"
	_ "crypto/md5"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	userDataTable  = "ws_chat.users"
	poolsDataTable = "ws_chat.pools"

	filesInfoTable = "file_server.files_info"
	hashCost       = 14
)

type PgSqlUserKeysManager struct {
	connection *sql.DB
}

func (db *PgSqlUserKeysManager) MakeConnection() {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	dbName := os.Getenv("POSTGRES_DB")

	connStr := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", user, password, host, port, dbName)
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	db.connection = conn
}

func (db *PgSqlUserKeysManager) CloseConnection() {
	if err := db.connection.Close(); err != nil {
		log.Error("Error closing fileInfoManager connection: ", err)
	}
}

func (db *PgSqlUserKeysManager) InitTables() {

	query := `create schema if not exists ` + strings.Split(userDataTable, ".")[0]

	if _, err := db.connection.Exec(query); err != nil {
		log.Fatal("Error creating schema: ", err)
	}

	query = `create table if not exists ` + userDataTable + `
(
    id            varchar(255)       not null
        constraint users_pk
            primary key,
    username      varchar(64)  not null,
    password      varchar(255) not null,
    email_address varchar(128) default NULL::character varying,
	pool_id		  varchar(255) default ''::character varying,
    m_token_secret   varchar(128) default ''::character varying,
	w_session_key varchar(255) default ''::character varying,
	user_secret varchar(512) not null,
	online		  boolean	   default false
)`
	if _, err := db.connection.Exec(query); err != nil {
		log.Fatal("Error creating table: ", err)
	}

	query = `create table if not exists ` + poolsDataTable + `
(
    id            serial       not null
        constraint pools_pk
            primary key,
    pool_id       varchar(255)  not null,
    password      varchar(255)  not null
)`
	if _, err := db.connection.Exec(query); err != nil {
		log.Fatal("Error creating table: ", err)
	}

	query = `create schema if not exists ` + strings.Split(filesInfoTable, ".")[0]
	if _, err := db.connection.Exec(query); err != nil {
		log.Fatal("Error creating table: ", err)
	}

	query = `create table if not exists ` + filesInfoTable + `
(
    lot_id            serial       not null
        constraint pools_pk
            primary key,
    file_id       varchar(255)  default null,
	chat_id			varchar(255)  not null,
    view_tokens      varchar(255)[]  not null default ARRAY[]::varchar(255)[],
	lot_status		integer			default 0,
	content_type	varchar(64) default null,
	file_size		integer			default 0
)`
	if _, err := db.connection.Exec(query); err != nil {
		log.Fatal("Error creating table: ", err)
	}
}

func (db *PgSqlUserKeysManager) SignInAndRefreshMobileAndWebKeys(user structs.User) (webKey string, mobileToken string, ok bool) {
	var truePassword string
	var id string

	query := `
SELECT id, password FROM ` + userDataTable + ` 
WHERE username = $1`

	row := db.connection.QueryRow(query, user.Username)
	if err := row.Scan(&id, &truePassword); err != nil {
		log.Errorf("Error getting user info: %v; user: %v", err, user.Username)
		return "", "", false
	}
	if err := bcrypt.CompareHashAndPassword([]byte(truePassword), []byte(user.Password)); err != nil {
		log.Error("Error verifying password: ", err)
		return "", "", false
	}
	sessionKeyBytes := md5.Sum([]byte(time.Now().String() + user.Username + strconv.Itoa(rand.Int())))
	cookieKey := fmt.Sprintf("%x", sessionKeyBytes)
	tokenSecret := hashes.CalculateSha1(user.Username + "_hey_" + time.Now().String() + string(rand.Int63()))

	if db.setUserSessionKeysById(id, cookieKey, tokenSecret) {
		return cookieKey, tokenSecret, true
	} else {
		return "", "", false
	}
}

func (db *PgSqlUserKeysManager) CreateUserAndGenerateKeys(user structs.User) (webKey string, mobileToken string, ok bool) {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), hashCost)
	if err != nil {
		log.Error("Error hashing password", err)
		return "", "", false
	}
	sessionKeyBytes := md5.Sum([]byte(time.Now().String() + user.Username + strconv.Itoa(rand.Int())))
	cookieKey := fmt.Sprintf("%x", sessionKeyBytes)
	tokenSecret := hashes.CalculateSha1(user.Username + "_hey_" + time.Now().String() + string(rand.Int63()))
	userSecret := hashes.CalculateSha256(strings.Join([]string{user.Username, user.Password, user.Username}, "&-"))

	query := `INSERT INTO ` + userDataTable + ` 
(id, username, password, email_address, m_token_secret, w_session_key, user_secret) 
VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err = db.connection.Exec(query, user.Id, user.Username, passwordHash, user.Email, tokenSecret, cookieKey, userSecret)

	if err != nil {
		log.Error("Error registering user: ", err)
		return "", "", false
	}
	return cookieKey, tokenSecret, true
}

func (db *PgSqlUserKeysManager) DeleteCookieKey(sessionKey string) {

	if !db.UpdateSessionKey(sessionKey, "") {
		log.Error("Error deleting session key")
	}
}

func (db *PgSqlUserKeysManager) DeleteUser(user structs.User) {
	query := `
DELETE FROM ` + userDataTable + ` 
WHERE username=$1`

	_, err := db.connection.Exec(query, user.Username)
	if err != nil {
		log.Error("Error deleting user: ", err)
	}
}

func (db *PgSqlUserKeysManager) ToggleUserOnlineState(id string, state bool) bool {
	query := `
UPDATE ` + userDataTable + ` 
SET online=$1 WHERE id=$2`

	if _, err := db.connection.Exec(query, state, id); err != nil {
		log.Error("Error updating online state: ", err)
		return false
	}
	log.Info("Toggled user %v state = %v", id, state)
	return true
}

func (db *PgSqlUserKeysManager) GetAllSamePoolUsers(id string) ([]structs.User, error) {

	var result []structs.User

	query := `SELECT username, online FROM ` + userDataTable + ` 
WHERE pool_id LIKE (
	SELECT pool_id FROM ` + userDataTable + `
	WHERE id=$1
)`
	rows, err := db.connection.Query(query, id)
	if err != nil {
		log.Error("Error getting users from same pool by id", err)
		return nil, err
	}
	for rows.Next() {
		userItem := &structs.User{}
		if err = rows.Scan(&userItem.Username, &userItem.IsOnline); err != nil {
			log.Error("Error scanning all users rows: ", err)
		} else {
			result = append(result, *userItem)
		}
	}
	return result, nil
}

func (db *PgSqlUserKeysManager) SetSessionKeyById(sessionKey string, id int) bool {
	query := `
UPDATE ` + userDataTable + ` 
SET session_key=$1
WHERE id=$2`

	if _, err := db.connection.Exec(query, sessionKey, id); err != nil {
		log.Error("Error setting session key: ", err)
		log.Error("Key: ", sessionKey, " ID: ", id)
		return false
	}
	return true
}

func (db *PgSqlUserKeysManager) SetUserSecret(user structs.User) bool {
	query := `
UPDATE ` + userDataTable + ` 
SET session_secret=$1
WHERE username=$2`

	if _, err := db.connection.Exec(query, user.Token, user.Username); err != nil {
		log.Error("Error setting session key: ", err)
		return false
	}
	return true
}

func (db *PgSqlUserKeysManager) UpdateSessionKey(old, new string) bool {
	query := `
UPDATE ` + userDataTable + ` 
SET session_key=$1
WHERE session_key=$2`

	if _, err := db.connection.Exec(query, new, old); err != nil {
		log.Error("Error updating session key: ", err)
		return false
	}
	return true
}

func (db *PgSqlUserKeysManager) CreatePool(pool structs.Pool) bool {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(pool.Password), hashCost)
	if err != nil {
		log.Error("Error hashing password", err)
		return false
	}

	query := `
INSERT INTO ` + poolsDataTable + `(pool_id, password) 
VALUES ($1, $2)`

	if _, err := db.connection.Exec(query, pool.PoolId, passwordHash); err != nil {
		log.Error("Error inserting pool: ", err)
		return false
	}
	return true
}

func (db *PgSqlUserKeysManager) TryPoolSignIn(pool structs.Pool) bool {
	var truePassword string

	query := `
SELECT password
FROM ` + poolsDataTable + `
WHERE pool_id=$1`

	row := db.connection.QueryRow(query, pool.PoolId)
	if err := row.Scan(&truePassword); err != nil {
		log.Error("Error getting pool password: ", err)
		return false
	}
	if err := bcrypt.CompareHashAndPassword([]byte(truePassword), []byte(pool.Password)); err != nil {
		log.Error("Error verifying password: ", err)
		return false
	}
	return true
}

func (db *PgSqlUserKeysManager) UpdateUserPoolId(id string, pool structs.Pool) bool {
	query := `
UPDATE ` + userDataTable + `
SET pool_id=$1
WHERE id=$2`

	_, err := db.connection.Exec(query, pool.PoolId, id)
	if err != nil {
		log.Error("error updating user pool id: ", err)
		return false
	}
	return true
}

func (db *PgSqlUserKeysManager) GetUserIdByName(name string) (string, bool) {
	var id string

	query := `
SELECT id
FROM ` + userDataTable + `
WHERE username=$1`

	row := db.connection.QueryRow(query, name)
	if err := row.Scan(&id); err != nil {
		log.Error("Error getting pool password: ", err)
		return "", false
	}
	return id, true
}

func (db *PgSqlUserKeysManager) GetUserNameById(id string) (string, bool) {
	var name string

	query := `
SELECT username
FROM ` + userDataTable + `
WHERE id=$1`

	row := db.connection.QueryRow(query, id)
	if err := row.Scan(&name); err != nil {
		log.Error("Error getting pool password: ", err)
		return "", false
	}
	return name, true
}

func (db *PgSqlUserKeysManager) CreateFileLot(chatId string) int64 {
	query := `
INSERT INTO ` + filesInfoTable + ` 
(chat_id) VALUES ($1) RETURNING lot_id`

	var id int64

	if err := db.connection.QueryRow(query, chatId).Scan(&id); err != nil {
		log.Errorf("Error creating file lot: %v", err)
		return -1
	}
	return id
}

func (db *PgSqlUserKeysManager) GetFileLotChatId(lotId string) string {
	query := `
SELECT chat_id FROM ` + filesInfoTable + ` 
WHERE lot_id=$1`

	var id string

	if err := db.connection.QueryRow(query, lotId).Scan(&id); err != nil {
		log.Errorf("Error creating file lot: %v", err)
		return ""
	}
	return id
}

func (db *PgSqlUserKeysManager) AddViewKeyToFileLot(lotId, key string) bool {
	query := `
UPDATE ` + filesInfoTable + ` 
SET view_tokens = array_append(view_tokens, $2)
WHERE lot_id=$1`

	if _, err := db.connection.Exec(query, lotId, key); err != nil {
		log.Errorf("Error creating file lot: %v", err)
		return false
	}
	return true
}
