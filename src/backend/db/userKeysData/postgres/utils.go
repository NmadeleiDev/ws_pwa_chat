package postgres

import "github.com/sirupsen/logrus"

func (db *PgSqlUserKeysManager) setUserSessionKeysById(id, cookieKey, tokenSecret string) bool {
	query := `
UPDATE ` + userDataTable + ` 
SET m_token_secret=$1, w_session_key=$2 
WHERE id=$3`

	if _, err := db.connection.Exec(query, tokenSecret, cookieKey, id); err != nil {
		logrus.Errorf("Error updating user keys: %v; id: %v", id)
		return false
	}
	return true
}