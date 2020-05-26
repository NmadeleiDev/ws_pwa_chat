package postgres

import (
	"chat_backend/structs"
	"crypto/md5"
	_ "crypto/md5"
	"database/sql"
	"errors"
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
	userDataTable = "ws_chat.users"
	poolsDataTable = "ws_chat.pools"
	hashCost = 14
)

var connection *sql.DB

func MakeConnection() {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	db := os.Getenv("POSTGRES_DB")

	connStr := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", user, password, host, port, db)
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	connection = conn
}

func CloseConnection() {
	if err := connection.Close(); err != nil {
		log.Error("Error closing postgres connection: ", err)
	}
}

func InitTables() {

	query := `create schema if not exists ` + strings.Split(userDataTable, ".")[0]

	if _, err := connection.Exec(query); err != nil {
		log.Fatal("Error creating schema: ", err)
	}

	query = `create table if not exists ` + userDataTable + `
(
    id            serial       not null
        constraint users_pk
            primary key,
    username      varchar(64)  not null,
    password      varchar(255) not null,
    email_address varchar(128) default NULL::character varying,
	pool_id		  varchar(255) default ''::character varying,
    session_key   varchar(128) default ''::character varying,
	session_secret varchar(255) default ''::character varying,
	online		  boolean	   default false
)`
	if _, err := connection.Exec(query); err != nil {
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
	if _, err := connection.Exec(query); err != nil {
		log.Fatal("Error creating table: ", err)
	}
}

func IssueUserSessionKey(user structs.User) (string, error) {
	var truePassword string
	var id int

	query := `
SELECT id, password FROM ` + userDataTable + ` 
WHERE username = $1`

	row := connection.QueryRow(query, user.Username)
	if err := row.Scan(&id, &truePassword); err != nil {
		log.Error("Error getting user info: ", err)
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(truePassword), []byte(user.Password)); err != nil {
		log.Error("Error verifying password: ", err)
		return "", err
	}
	sessionKeyBytes := md5.Sum([]byte(time.Now().String() + user.Username + strconv.Itoa(rand.Int())))
	sessionKey := fmt.Sprintf("%x", sessionKeyBytes)

	if SetSessionKeyById(sessionKey, id) {
		return sessionKey, nil
	} else {
		return "", errors.New("error updating session key")
	}
}

func CreateUser(user structs.User) bool {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), hashCost)
	if err != nil {
		log.Error("Error hashing password", err)
		return false
	}

	query := `INSERT INTO ` + userDataTable + ` (username, password, email_address, session_secret) VALUES ($1, $2, $3, $4)`
	_, err = connection.Exec(query, user.Username, passwordHash, user.Email, user.SecretHash)

	if err != nil {
		log.Error("Error registering user", err)
		return false
	}
	return true
}

func DeleteSessionKey(sessionKey string) {

	if !UpdateSessionKey(sessionKey, "") {
		log.Error("Error deleting session key")
	}
}

func DeleteUser(user structs.User)  {
	query := `
DELETE FROM ` + userDataTable + ` 
WHERE username=$1`

	_, err := connection.Exec(query, user.Username)
	if err != nil {
		log.Error("Error deleting user: ", err)
	}
}

func GetUserNameIdAndPool(sessionKey string) (user structs.User, err error) {

	query := `
SELECT username, id, pool_id
FROM ` + userDataTable + ` 
WHERE session_key=$1`

	row := connection.QueryRow(query, sessionKey)
	err = row.Scan(&user.Username, &user.Id, &user.Pool)
	return user, err
}

func ToggleUserOnlineState(id int, state bool) bool {
	query := `
UPDATE ` + userDataTable + ` 
SET online=$1 WHERE id=$2`

	if _, err := connection.Exec(query, state, id); err != nil {
		log.Error("Error updating online state: ", err)
		return false
	}
	log.Info("Toggled user %v state = %v", id, state)
	return true
}

func GetAllSamePoolUsers(userPool string) ([]structs.User, error) {

	var result []structs.User

	query := `SELECT username FROM ` + userDataTable + ` WHERE pool_id=$1`
	rows, err := connection.Query(query, userPool)
	if err != nil {
		log.Error("Error querying users", err)
		return nil, err
	}
	for rows.Next() {
		userItem := &structs.User{}
		if err = rows.Scan(&userItem.Username); err != nil {
			log.Error("Error scanning all users rows: ", err)
		} else {
			result = append(result, *userItem)
		}
	}
	return result, nil
}

func SetSessionKeyById(sessionKey string, id int) bool {
	query := `
UPDATE ` + userDataTable + ` 
SET session_key=$1
WHERE id=$2`

	if _, err := connection.Exec(query, sessionKey, id); err != nil {
		log.Error("Error setting session key: ", err)
		log.Error("Key: ", sessionKey, " ID: ", id)
		return false
	}
	return true
}

func SetUserSecret(user structs.User) bool {
	query := `
UPDATE ` + userDataTable + ` 
SET session_secret=$1
WHERE username=$2`

	if _, err := connection.Exec(query, user.SecretHash, user.Username); err != nil {
		log.Error("Error setting session key: ", err)
		return false
	}
	return true
}

func UpdateSessionKey(old, new string) bool {
	query := `
UPDATE ` + userDataTable + ` 
SET session_key=$1
WHERE session_key=$2`

	if _, err := connection.Exec(query, new, old); err != nil {
		log.Error("Error updating session key: ",err)
		return false
	}
	return true
}

func CreatePool(pool structs.Pool) bool {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(pool.Password), hashCost)
	if err != nil {
		log.Error("Error hashing password", err)
		return false
	}

	query := `
INSERT INTO ` + poolsDataTable + `(pool_id, password) 
VALUES ($1, $2)`

	if _, err := connection.Exec(query, pool.PoolId, passwordHash); err != nil {
		log.Error("Error inserting pool: ",err)
		return false
	}
	return true
}

func TryPoolSignIn(pool structs.Pool) bool {
	var truePassword string

	query := `
SELECT password
FROM ` + poolsDataTable + `
WHERE pool_id=$1`

	row := connection.QueryRow(query, pool.PoolId)
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

func UpdateUserPoolId(user structs.User, pool structs.Pool) bool {
	query := `
UPDATE ` + userDataTable + `
SET pool_id=$1
WHERE username=$2`

	_, err := connection.Exec(query, pool.PoolId, user.Username)
	if err != nil {
		log.Error("error updating user pool id: ", err)
		return false
	}
	return true
}