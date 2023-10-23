package datamanagement

type employeeFromDb struct {
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
}

type workplaceFromDb struct {
	IdWorkplace int
	Name        string
	Adress      string
}

type sitesPageStruct struct {
	Workplace []workplaceFromDb
	Employee  []employeeFromDb
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