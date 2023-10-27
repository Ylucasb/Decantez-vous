package datamanagement

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
	BirthDate    string
	HireDate     string
	IBAN         string
	IsWorking    bool
	IsPays       bool
	Job          string
}

type WorkplaceFromDb struct {
	IdWorkplace int
	Name        string
	Adress      string
	Phone       string
	Mail        string
	Type        string
}

type SupplierFromDb struct {
	IdSupplier int
	IdProduct  int
	FirstName  string
	LastName   string
	Adress     string
	Phone      string
	Mail       string
}
