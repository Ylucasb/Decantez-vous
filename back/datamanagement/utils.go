package datamanagement

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func SelectDB(query string, args ...interface{}) *sql.Rows {
	db, err := sql.Open("sqlite3", "./database/decantez-vous.db")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer db.Close()
	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return rows
}

func AddDeleteUpdateDB(query string, args ...interface{}) sql.Result {
	db, err := sql.Open("sqlite3", "./database/decantez-vous.db")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer db.Close()

	res, err := db.Exec(query, args...)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return res
}

func GetCookieValue(cookie *http.Cookie) string {
	var valueReturned string
	test := false
	value := fmt.Sprint(cookie)
	for _, element := range value {
		if test {
			valueReturned += string(element)
		}
		if element == 61 {
			test = true
		}
	}
	return valueReturned
}

func Hash(password string) string {
	passwordByte := []byte("decantez-vous" + password + "decantez-vous")
	passwordInSha256 := sha256.Sum256(passwordByte)
	stringPasswordInSha256 := fmt.Sprintf("%x", passwordInSha256[:])
	return stringPasswordInSha256
}
