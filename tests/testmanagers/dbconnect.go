package testmanagers

import (
	"os"

	dbConnManager "github.com/udistrital/cuentas_contables_crud/db"
)

// MakeDbConnection connection test
func MakeDbConnection() {
	dbConnManager.InitDB(
		os.Getenv("CUENTAS_CONTABLES_CRUD_DB_URL"),
		"27017",
		os.Getenv("CUENTAS_CONTABLES_CRUD_DB_USER"),
		os.Getenv("CUENTAS_CONTABLES_CRUD_DB_PASS"),
		os.Getenv("CUENTAS_CONTABLES_CRUD_DB_AUTH"),
		os.Getenv("CUENTAS_CONTABLES_CRUD_DB_NAME"),
	)

}
