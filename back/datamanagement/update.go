package datamanagement

func Work(idUser int) {
	AddDeleteUpdateDB("UPDATE employee SET isWorking=not isWorking WHERE idEmployee=?;", idUser)
}
