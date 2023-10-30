package datamanagement

func Work(idUser int) {
	AddDeleteUpdateDB("UPDATE employee SET isWorking=not isWorking WHERE idEmployee=?;", idUser)
}

func UpdateIban(idUser int, newIban string) {
	AddDeleteUpdateDB("UPDATE employee SET IBAN=? WHERE idEmployee=?;", newIban, idUser)
}
