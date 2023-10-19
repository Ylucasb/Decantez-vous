package datamanagement

import (
	"database/sql"
	"fmt"
	"net/http"
)

func SelectDB(query string, args ...interface{}) *sql.Rows {
	db, err := sql.Open("sqlite3", "./database/décantez-vous.db")
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
