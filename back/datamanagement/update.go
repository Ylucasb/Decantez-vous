package datamanagement

import "strconv"

func Work(idUser int) {
	AddDeleteUpdateDB("UPDATE employee SET isWorking=not isWorking WHERE idEmployee=?;", idUser)
}

func UpdateIban(idUser int, newIban string) {
	AddDeleteUpdateDB("UPDATE employee SET IBAN=? WHERE idEmployee=?;", newIban, idUser)
}
func ChangeWork(changeWork []string) {
	intIdWork, _ := strconv.Atoi(changeWork[0])
	intIdEmployee, _ := strconv.Atoi(changeWork[1])
	AddDeleteUpdateDB("UPDATE employee SET idProfession = ? WHERE idEmployee=?;", intIdWork, intIdEmployee)
}
