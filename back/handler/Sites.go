package handler

import (
	"decantez-vous/back/datamanagement"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type structDisplaySites struct {
	Employee     []datamanagement.EmployeeFromDb
	Workplace    []datamanagement.WorkplaceFromDb
	IsPays       bool
	ErrorGestion string
}

func Disconnect(w http.ResponseWriter, r *http.Request) {
	cookieIsConnected := http.Cookie{Name: "isConnected", Value: "", Expires: time.Now()}
	http.SetCookie(w, &cookieIsConnected)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func isValidDateFormat(dateString string) bool {
	_, err := time.Parse("02-01-2006", dateString)
	return err == nil
}

func Sites(w http.ResponseWriter, r *http.Request, adress string) {
	t := template.Must(template.ParseFiles("./front/html/sites.html"))
	cookieIsPays, _ := r.Cookie("isPays")
	cookieIdUser, _ := r.Cookie("idUser")
	isPays := datamanagement.GetCookieValue(cookieIsPays)
	idUser := datamanagement.GetCookieValue(cookieIdUser)
	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	phone := r.FormValue("phone")
	mail := r.FormValue("mail")
	password := r.FormValue("password")
	IBAN := r.FormValue("IBAN")
	birthDate := r.FormValue("birthDate")
	idProfession := r.FormValue("profession")
	disconnect := r.FormValue("disconnect")
	work := r.FormValue("work")
	deleteEmployee := r.FormValue("deleteEmployee")
	changeWork := r.FormValue("changeWork")
	nIban := r.FormValue("nIban")
	idUserSelected := r.FormValue("idUserSelected")
	errorGestion := ""

	//change IBAN
	if nIban != "" && idUserSelected != "" {
		idUser, err := strconv.Atoi(idUserSelected)
		if err != nil {
			log.Fatal(err)
		}
		datamanagement.UpdateIban(idUser, nIban)
	}
	//changeWork
	if changeWork != "" {
		changeWork := strings.Split(changeWork, ",")
		datamanagement.ChangeWork(changeWork)
	}
	// delete employee
	if deleteEmployee != "" {
		idDeleteEmployee, err := strconv.Atoi(deleteEmployee)
		if err != nil {
			log.Fatal(err)
		} else {
			datamanagement.DeleteEmployee(idDeleteEmployee)
		}
	}
	// Disconnect user
	if disconnect != "" {
		Disconnect(w, r)
	}
	if work != "" {
		intIdUser, err := strconv.Atoi(idUser)
		if err == nil {
			datamanagement.Work(intIdUser)
		}
	}

	if firstName != "" && lastName != "" && phone != "" && mail != "" && password != "" && IBAN != "" && birthDate != "" && idProfession != "" {
		intIdUser, err := strconv.Atoi(idUser)
		intIdProfession, err := strconv.Atoi(idProfession)
		if err == nil && isValidDateFormat(birthDate) {
			date, _ := time.Parse("02-01-2006", birthDate)
			datamanagement.AddEmployee(firstName, lastName, phone, mail, password, date.Format("2006-01-02"), IBAN, adress, intIdProfession, intIdUser)
		} else {
			errorGestion = "erreur a la creation de l'employé"
		}
	}

	var isPaysBool bool
	if isPays == "true" {
		isPaysBool = true
	} else {
		isPaysBool = false
	}
	structDisplaySites := structDisplaySites{datamanagement.RecuperationEmployeeWorkplace(adress), datamanagement.RecuperationWorkplace(adress), isPaysBool, errorGestion}
	if isPaysBool {
		for i := 0; i < len(structDisplaySites.Employee); i++ {
			structDisplaySites.Employee[i].IsPays = isPaysBool
		}
	}
	err := t.Execute(w, structDisplaySites)
	if err != nil {
		log.Fatal(err)
	}
}
