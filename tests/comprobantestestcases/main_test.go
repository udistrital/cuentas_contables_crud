package comprobantestestcases

import (
	"os"
	"testing"

	"github.com/udistrital/cuentas_contables_crud/tests/testmanagers"
)

func TestMain(m *testing.M) {
	testmanagers.MakeDbConnection()
	exitVal := m.Run()
	os.Exit(exitVal)
}
