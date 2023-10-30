package datamanagement

import (
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func AddSupplier(firstName string, lastName string, product string, phone string, adress string, mail string, workplace []string) {
	idProduct := RecuperationProduct(product)
	AddDeleteUpdateDB("INSERT INTO supplier (firstName, lastName, idProduct, phone, adress, mail) VALUES (?, ?, ?, ?, ?, ?);", firstName, lastName, idProduct, phone, adress, mail)
	idSupplier := RecuperationIdSupplier(firstName, lastName, phone, adress, mail)
	for i := 0; i < len(workplace); i++ {
		idWorkplace := RecuperationWorkplace(workplace[i])
		AddDeleteUpdateDB("INSERT INTO relationWorkplaceSupplier (idSupplier, idWorkplace ) VALUES (?, ?);", idSupplier, idWorkplace[0].IdWorkplace)
	}
}

func AddEmployee(firstName string, lastName string, phone string, mail string, password string, birthDate string, IBAN string, workPlaceName string, profession int, superior int) {
	AddDeleteUpdateDB("INSERT INTO employee (firstName, lastName, phone, mail, password, birthDate, hireDate, iban, idWorkplace, isWorking, idProfession) VALUES (?,?,?,?,?,?,?,?,(SELECT idWorkplace FROM workplace WHERE name =?),?,?);", firstName, lastName, phone, mail, Hash(password), birthDate, time.Now(), IBAN, workPlaceName, false, profession)
	AddDeleteUpdateDB("INSERT INTO relationEmployee (idEmployee, idReferent) VALUES ((SELECT idEmployee FROM employee WHERE mail = ?),(SELECT idEmployee FROM employee WHERE idEmployee = ?));", mail, superior)
}
