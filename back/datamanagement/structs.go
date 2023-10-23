package datamanagement

import "time"

type EmployeeFromDb struct {
	IdEmployee   int
	IdWorkplace  int
	IdRelation   int
	IdProfession int
	FirstName    string
	LastName     string
	Phone        string
	Mail         string
	Password     string
	BirthDate    time.Time
	HireDate     time.Time
	IBAN         string
	IsWorking    bool
	IsPays       bool
	Job          string
}

type workplaceFromDb struct {
	IdWorkplace int
	Name        string
	Adress      string
}

type sitesPageStruct struct {
	Workplace []workplaceFromDb
	Employee  []EmployeeFromDb
	// NbrPost     int
	IsConnected bool
}

type supplierFromDb struct {
	IdSupplier int
	IdProduct  int
	FirstName  string
	LastName   string
	Adress     string
	Phone      string
	Mail       string
}
