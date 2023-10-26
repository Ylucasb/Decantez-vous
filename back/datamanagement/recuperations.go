package datamanagement

import (
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func getJob(allEmployees []EmployeeFromDb) []EmployeeFromDb {
	rows := SelectDB("SELECT idProfession,name FROM profession")
	saveJob := map[int]string{}
	defer rows.Close()
	for rows.Next() {
		var idProfession int
		var name string
		err := rows.Scan(&idProfession, &name)
		if err != nil {
			log.Fatal(err)
		}
		saveJob[idProfession] = name
	}
	for i := 0; i < len(allEmployees); i++ {
		allEmployees[i].Job = saveJob[allEmployees[i].IdProfession]
	}
	return allEmployees
}

func RecuperationEmployeeWorkplace(workplace string) []EmployeeFromDb {
	var allEmployees []EmployeeFromDb
	rows := SelectDB("SELECT * FROM employee INNER JOIN workplace ON workplace.idWorkplace = employee.idWorkplace WHERE workplace.adress = ?", workplace)
	defer rows.Close()
	for rows.Next() {
		var idEmployee int
		var idWorkplace int
		var idRelation int
		var idProfession int
		var firstName string
		var lastName string
		var phone string
		var mail string
		var password string
		var birthDate time.Time
		var hireDate time.Time
		var iban string
		var isWorking bool
		var idWorkplace2 int
		var name string
		var adress string
		err := rows.Scan(&idEmployee, &firstName, &lastName, &phone, &mail, &password, &birthDate, &hireDate, &iban, &idRelation, &idWorkplace, &isWorking, &idProfession, &idWorkplace2, &name, &adress)
		if err != nil {
			log.Fatal(err)
		}
		employeeStruc := EmployeeFromDb{
			IdEmployee:   idEmployee,
			IdWorkplace:  idWorkplace,
			IdRelation:   idRelation,
			IdProfession: idProfession,
			FirstName:    firstName,
			LastName:     lastName,
			Phone:        phone,
			Mail:         mail,
			Password:     password,
			BirthDate:    birthDate.Format("02-01-2006"),
			HireDate:     hireDate.Format("02-01-2006"),
			IBAN:         iban,
			IsWorking:    isWorking,
			IsPays:       false,
			Job:          "",
		}
		allEmployees = append(allEmployees, employeeStruc)
	}
	allEmployees = getJob(allEmployees)
	return allEmployees
}

func RecuperationSupplier() []SupplierFromDb {
	var allSuppliers []SupplierFromDb
	rows := SelectDB("SELECT IdSupplier,product.name,FirstName,LastName,Adress,Phone,Mail FROM supplier INNER JOIN product ON product.idProduct = supplier.idProduct")
	defer rows.Close()
	for rows.Next() {
		var idSupplier int
		var product string
		var firstName string
		var lastName string
		var adress string
		var phone string
		var mail string

		err := rows.Scan(&idSupplier, &product, &firstName, &lastName, &adress, &phone, &mail)
		if err != nil {
			log.Fatal(err)
		}
		supplierStruc := SupplierFromDb{
			IdSupplier: idSupplier,
			Product:    product,
			FirstName:  firstName,
			LastName:   lastName,
			Adress:     adress,
			Phone:      phone,
			Mail:       mail,
		}
		allSuppliers = append(allSuppliers, supplierStruc)
	}

	return allSuppliers
}

func RecuperationSupplierWorkplace() []SupplierWorkplaceFromDb {
	var allSuppliersWorkplace []SupplierWorkplaceFromDb
	rows := SelectDB("SELECT workplace.name, idSupplier FROM relationWorkplaceSupplier INNER JOIN workplace ON workplace.idWorkplace = relationWorkplaceSupplier.idWorkplace")
	defer rows.Close()
	for rows.Next() {
		var workplace string
		var idSupplier int

		err := rows.Scan(&workplace, &idSupplier)
		if err != nil {
			log.Fatal(err)
		}
		supplierWorkplaceStruc := SupplierWorkplaceFromDb{
			Workplace:  workplace,
			IdSupplier: idSupplier,
		}
		allSuppliersWorkplace = append(allSuppliersWorkplace, supplierWorkplaceStruc)
	}

	return allSuppliersWorkplace
}
