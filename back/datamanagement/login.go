package datamanagement

import (
	"crypto/sha256"
	"fmt"
)

func IsRegister(userInput string, password string) (bool, string) {
	passwordByte := []byte(password)
	passwordInSha256 := sha256.Sum256(passwordByte)
	stringPasswordInSha256 := fmt.Sprintf("%x", passwordInSha256[:])
	rows := SelectDB("SELECT userID FROM Users WHERE (Email = ? ) AND Password = ?;", string(userInput), string(stringPasswordInSha256))
	defer rows.Close()
	for rows.Next() {
		var id string
		rows.Scan(&id)
		return true, id
	}
	return false, ""
}
