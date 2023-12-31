package datamanagement

import (
	_ "github.com/mattn/go-sqlite3"
)

func DeleteSupplier(idDeleteSupplier int) {
	AddDeleteUpdateDB("DELETE FROM supplier WHERE idSupplier = ?", idDeleteSupplier)
	AddDeleteUpdateDB("DELETE FROM relationWorkplaceSupplier WHERE idSupplier = ?", idDeleteSupplier)
}

func DeleteEmployee(idEmployee int) {
	AddDeleteUpdateDB("DELETE FROM employee WHERE idEmployee = ?;", idEmployee)
}
