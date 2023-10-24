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
