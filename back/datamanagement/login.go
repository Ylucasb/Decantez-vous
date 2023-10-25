package datamanagement

import (
	"crypto/sha256"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func IsRegister(userInput string, password string) (bool, int, bool) {
	var isExist bool = false
	var isPays bool = false
	var id int = 0
	passwordByte := []byte("decantez-vous" + password + "decantez-vous")
	passwordInSha256 := sha256.Sum256(passwordByte)
	stringPasswordInSha256 := fmt.Sprintf("%x", passwordInSha256[:])
	rows := SelectDB("SELECT idEmployee,idProfession FROM employee WHERE mail = ? AND password = ?;", string(userInput), string(stringPasswordInSha256))
	defer rows.Close()
	for rows.Next() {
		var idE int
		var idP int
		rows.Scan(&idE, &idP)
		isExist = true
		id = idE
		if idP < 3 {
			isPays = true
		}
	}
	return isExist, id, isPays
}

func TestIsConnected(r *http.Request) bool {
	cookieConnected, _ := r.Cookie("isConnected")
	IsConnected := GetCookieValue(cookieConnected)
	if IsConnected == "true" {
		return true
	} else {
		return false
	}
}
