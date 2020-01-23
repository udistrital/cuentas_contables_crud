package testmanagers

import (
	dbConnManager "github.com/udistrital/cuentas_contables_crud/db"
	"testing"
)

func TestDbConnectSuccess(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Error("error: ", r)
			t.Fail()
		} else {
			t.Log("TestDbConnectSuccess Finalizado Correctamente (OK)")

		}
	}()

	if _, err := dbConnManager.GetDatabase(); err != nil {
		panic(err.Error())
	}

}
