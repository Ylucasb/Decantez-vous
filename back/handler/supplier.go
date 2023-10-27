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
	Supplier []datamanagement.SupplierFromDb
}

func Supplier(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./front/html/supplier.html"))

	cookieIsPays, _ := r.Cookie("isPays")
	isPays := datamanagement.GetCookieValue(cookieIsPays)
	var isPaysBool bool
	if isPays == "true" {
		isPaysBool = true
	} else {
		isPaysBool = false
	}

	structDisplaySupplier := structDisplaySupplier{datamanagement.RecuperationSupplier()}

	deleteSupplier := r.FormValue("deleteSupplier")
	if deleteSupplier != "" {
		i, err := strconv.Atoi(deleteSupplier)
		if err != nil {
			log.Fatal(err)
		}
		idDeleteSupplier := structDisplaySupplier.Supplier[i-1].IdSupplier
		datamanagement.DeleteSupplier(idDeleteSupplier)
	}

	structDisplaySupplierAfterModif := structDisplaySupplierAfterModif{datamanagement.RecuperationSupplier()}
	allSuppliersWorkplace := datamanagement.RecuperationSupplierWorkplace()

	if isPaysBool {
		for i := 0; i < len(structDisplaySupplierAfterModif.Supplier); i++ {
			structDisplaySupplierAfterModif.Supplier[i].IsPays = isPaysBool
		}
	}

	for i := 0; i < len(structDisplaySupplierAfterModif.Supplier)-1; i++ {
		sameSupplier := []string{}
		sameSupplier = nil
		for j := 0; j < len(allSuppliersWorkplace)-1; j++ {
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
