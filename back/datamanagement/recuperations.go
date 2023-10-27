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
	rows := SelectDB("SELECT employee.idEmployee, employee.firstName, employee.lastName, employee.phone, employee.mail, employee.password, employee.birthDate, employee.hireDate, employee.IBAN, employee.idRelation, employee.idWorkplace , employee.isWorking, employee.idProfession FROM employee INNER JOIN workplace ON workplace.idWorkplace = employee.idWorkplace WHERE workplace.name = ?", string(workplace))
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

func RecuperationWorkplace(workplaceAdress string) []WorkplaceFromDb {
	var workplace []WorkplaceFromDb
	rows := SelectDB("SELECT * FROM workplace WHERE workplace.name = ?", workplaceAdress)
	defer rows.Close()
	for rows.Next() {
		var idWorkplace int
		var name string
		var adress string
		var phone string
		var mail string
		var typeWorkplace string
		err := rows.Scan(&idWorkplace, &name, &adress, &phone, &mail, &typeWorkplace)
		if err != nil {
			log.Fatal(err)
		}
		workplaceStruc := WorkplaceFromDb{
			IdWorkplace: idWorkplace,
			Name:        name,
			Adress:      adress,
			Phone:       phone,
			Mail:        mail,
			Type:        typeWorkplace,
		}
		workplace = append(workplace, workplaceStruc)
	}
	return workplace
}
