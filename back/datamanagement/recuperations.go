package datamanagement

import (
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func RecuperationEmployee() []EmployeeFromDb {
	var allEmployees []EmployeeFromDb
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
		var birthDate time.Time
		var hireDate time.Time
		var iban string
		var isWorking bool
		err := rows.Scan(&idEmployee, &firstName, &lastName, &phone, &mail, &password, &birthDate, &hireDate, &iban, &idRelation, &idWorkplace, &isWorking, &idProfession)
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
			BirthDate:    birthDate,
			HireDate:     hireDate,
			IBAN:         iban,
			IsWorking:    isWorking,
			IsPays:       false,
			Job:          "",
		}
		if employeeStruc.IdProfession < 3 {
			employeeStruc.IsPays = true
		}
		allEmployees = append(allEmployees, employeeStruc)
	}
	allEmployees = getJob(allEmployees)
	return allEmployees
}

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
