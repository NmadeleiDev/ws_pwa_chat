package postgres

import (
	"github.com/sirupsen/logrus"
)

func (db *PgSqlUserKeysManager) IdentifyUserByCookie(cookie string) (id string, ok bool) {
	query := `
SELECT id FROM ` + userDataTable + ` 
WHERE w_session_key=$1`

	if err := db.connection.QueryRow(query, cookie).Scan(&id); err != nil {
		logrus.Errorf("Error scanning id from cookie: %v; id: %v", err, id)
		return "", false
	}

	logrus.Infof("identified user by cookie. Got Id: %v;", id)

	return id, true
}

//func (db *PgSqlUserKeysManager) IdentifyUserByCookie(cookie string) (id, newCookie string, ok bool) {
//	query := `
//SELECT id FROM ` + userDataTable + `
//WHERE w_session_key=$1`
//
//	if err := db.connection.QueryRow(query, cookie).Scan(&id); err != nil {
//		logrus.Errorf("Error scanning id from cookie: %v; id: %v", err, id)
//		return "", "", false
//	}
//
//	newCookie = fmt.Sprintf("%v", md5.Sum([]byte(strings.Join([]string{cookie, time.Now().String(), string(rand.Int63())}, "oop"))))
//	logrus.Infof("identified user by cookie. Got Id: %v; Setting new cookie: %v", id, newCookie)
//
//	query = `
//UPDATE ` + userDataTable + `
//SET w_session_key=$1
//WHERE id=$2`
//
//	if _, err := db.connection.Exec(query, newCookie, id); err != nil {
//		logrus.Errorf("Error updating cookie: %v; id: %v", err, id)
//		return "", "", false
//	}
//	return id, newCookie, ok
//}
