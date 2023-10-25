package handler

import (
	"decantez-vous/back/datamanagement"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type structDisplaySupplier struct {
	Supplier []datamanagement.SupplierFromDb
}

func Supplier(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./front/html/supplier.html"))

	structDisplaySupplier := structDisplaySupplier{datamanagement.RecuperationSupplier()}
	allSuppliersWorkplace := datamanagement.RecuperationSupplierWorkplace()

	for i := 0; i < len(structDisplaySupplier.Supplier)-1; i++ {

		sameSupplier := []string{}
		sameSupplier = nil
		for j := 0; j < len(allSuppliersWorkplace)-1; j++ {
			if structDisplaySupplier.Supplier[i].IdSupplier == allSuppliersWorkplace[j].IdSupplier {
				sameSupplier = append(sameSupplier, allSuppliersWorkplace[j].Workplace)
			}
		}
		structDisplaySupplier.Supplier[i].Workplace = sameSupplier
	}

	err := t.Execute(w, structDisplaySupplier)
	if err != nil {
		log.Fatal(err)
	}
}
