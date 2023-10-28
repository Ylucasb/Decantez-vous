package datamanagement

import (
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
