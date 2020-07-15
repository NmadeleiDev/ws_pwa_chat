package postgres

import (
	"chat_backend/hashes"
	"chat_backend/structs"
	"github.com/sirupsen/logrus"
	"strings"
)

func (db *PgSqlUserKeysManager) GetUserKeysByName(username string) (secretKey, userSecret string, ok bool) {

	query := `
SELECT session_secret, user_secret FROM ` + userDataTable + ` 
WHERE username=$1`

	if err := db.connection.QueryRow(query, username).Scan(&secretKey, &userSecret); err != nil {
		logrus.Errorf("Error getting user keys: %v; name: %v", err, username)
		return "", "", false
	}
	return secretKey, userSecret, true
}

func (db *PgSqlUserKeysManager) IdentifyUserByToken(auth structs.MobileToken, time string) (id string, ok bool) {
	var tokenSecret string
	var userSecret string

	query := `
SELECT id, m_token_secret, user_secret FROM ` + userDataTable + ` 
WHERE username=$1`

	row := db.connection.QueryRow(query, auth.Username)
	if err := row.Scan(&id, &tokenSecret, &userSecret); err != nil {
		logrus.Errorf("Error scanning user id: %v; auth: %v; time: %v", err, auth, time)
		return "", false
	}

	token := hashes.CalculateSha256(strings.Join([]string{auth.Username, time, tokenSecret, userSecret}, ""))
	if token == auth.Token {
		return id, true
	} else {
		return "", false
	}
}
