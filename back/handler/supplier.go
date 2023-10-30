package handler

import (
	"decantez-vous/back/datamanagement"
	"html/template"
	"log"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type structDisplaySupplier struct {
	Supplier []datamanagement.SupplierFromDb
}

type structDisplaySupplierAfterModif struct {
	Supplier      []datamanagement.SupplierFromDb
	CanAddSuplier bool
}

func Supplier(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./front/html/supplier.html"))
	cookieIsPays, _ := r.Cookie("canAddSuplier")
	canAddSuplier := datamanagement.GetCookieValue(cookieIsPays)
	var canAddSuplierBool bool
	if canAddSuplier == "true" {
		canAddSuplierBool = true
	} else {
		canAddSuplierBool = false
	}

	//add supplier
	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	product := r.FormValue("product")
	phone := r.FormValue("phone")
	adress := r.FormValue("adress")
	mail := r.FormValue("mail")
	Nantes := r.FormValue("nantesWorkplace")
	Marseille := r.FormValue("marseilleWorkplace")
	Lille := r.FormValue("lilleWorkplace")
	Bordeaux := r.FormValue("bordeauxWorkplace")
	Strasbourg := r.FormValue("strasbourgWorkplace")
	Lyon := r.FormValue("lyonWorkplace")

	//slice for workplace
	var addSupplierWorkplace []string
	if Nantes != "" {
		addSupplierWorkplace = append(addSupplierWorkplace, Nantes)
	}
	if Marseille != "" {
		addSupplierWorkplace = append(addSupplierWorkplace, Marseille)
	}
	if Lille != "" {
		addSupplierWorkplace = append(addSupplierWorkplace, Lille)
	}
	if Bordeaux != "" {
		addSupplierWorkplace = append(addSupplierWorkplace, Bordeaux)
	}
	if Strasbourg != "" {
		addSupplierWorkplace = append(addSupplierWorkplace, Strasbourg)
	}
	if Lyon != "" {
		addSupplierWorkplace = append(addSupplierWorkplace, Lyon)
	}
	if firstName != "" && lastName != "" && product != "" && phone != "" && adress != "" && mail != "" {
		datamanagement.AddSupplier(firstName, lastName, product, phone, adress, mail, addSupplierWorkplace)
	}

	// delete supplier
	deleteSupplier := r.FormValue("deleteSupplier")
	if deleteSupplier != "" {
		idDeleteSupplier, err := strconv.Atoi(deleteSupplier)
		if err != nil {
			log.Fatal(err)
		}
		datamanagement.DeleteSupplier(idDeleteSupplier)
	}

	structDisplaySupplierAfterModif := structDisplaySupplierAfterModif{datamanagement.RecuperationSupplier(), canAddSuplierBool}
	allSuppliersWorkplace := datamanagement.RecuperationSupplierWorkplace()

	if canAddSuplierBool {
		for i := 0; i < len(structDisplaySupplierAfterModif.Supplier); i++ {
			structDisplaySupplierAfterModif.Supplier[i].DeleteAuthorization = canAddSuplierBool
		}
	}

	// add workplace to the structure
	for i := 0; i < len(structDisplaySupplierAfterModif.Supplier); i++ {
		sameSupplier := []string{}
		sameSupplier = nil
		for j := 0; j < len(allSuppliersWorkplace); j++ {
			if structDisplaySupplierAfterModif.Supplier[i].IdSupplier == allSuppliersWorkplace[j].IdSupplier {
				sameSupplier = append(sameSupplier, allSuppliersWorkplace[j].Workplace)
			}
		}
		structDisplaySupplierAfterModif.Supplier[i].Workplace = sameSupplier
	}

	err := t.Execute(w, structDisplaySupplierAfterModif)
	if err != nil {
		log.Fatal(err)
	}
}
