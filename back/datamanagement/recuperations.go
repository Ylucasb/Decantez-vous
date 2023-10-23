package datamanagement

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var allEmployees []employeeFromDb

func RecuperationEmployee() []employeeFromDb {
	rows := SelectDB("SELECT * FROM employee")
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
		var birthDate string
		var hireDate string
		var iban string
		var isWorking bool
		err := rows.Scan(&idEmployee, &idWorkplace, &idRelation, &idProfession, &firstName, &lastName, &phone, &mail, &password, &birthDate, &hireDate, &iban, &isWorking)
		if err != nil {
			log.Fatal(err)
		}
		employeeStruc := employeeFromDb{}
		employeeStruc = employeeFromDb{
			IdEmployee:   idEmployee,
			IdWorkplace:  idWorkplace,
			IdRelation:   idRelation,
			IdProfession: idProfession,
			FirstName:    firstName,
			LastName:     lastName,
			Phone:        phone,
			Mail:         mail,
			Password:     password,
			BirthDate:    birthDate,
			HireDate:     hireDate,
			IBAN:         iban,
			IsWorking:    isWorking,
		}
		allEmployees = append(allEmployees, employeeStruc)
	}
	return allEmployees
}

var allWorplace []workplaceFromDb

func RecuperationWorkplace() []workplaceFromDb {
	rows := SelectDB("SELECT * FROM workplace")
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var adress string
		err := rows.Scan(&id, &name, &adress)
		if err != nil {
			log.Fatal(err)
		}
		workplaceStruc := workplaceFromDb{}
		workplaceStruc = workplaceFromDb{
			IdWorkplace: id,
			Name:        name,
			Adress:      adress,
		}
		allWorplace = append(allWorplace, workplaceStruc)
	}
	return allWorplace
}
