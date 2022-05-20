package nodoCuentaContableTestCases

import (
	"testing"

	"github.com/udistrital/cuentas_contables_crud/managers"
	"github.com/udistrital/cuentas_contables_crud/models"
)

func TestNodoFail(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Error("error: ", r)
			t.Fail()
		} else {
			t.Log("TestFailNode Finalizado Correctamente (OK)")
		}
	}()

	nodoToInsert := models.NodoCuentaContable{
		Codigo: "1",
	}

	mang := managers.NodoCuentaContableManager{
		// Ctx: ctx, // set this bar if mongo is deployed on replica set mode.
	}
	if err := mang.AddNode(&nodoToInsert); err == nil {
		panic(err.Error())
	}

}

// func TestNodoFail(t *testing.T) {

// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Log("TestComprobanteFail Finalizado Correctamente (OK)")
// 		} else {
// 			t.Error("error: Comprobante doesn't create ")
// 			t.Fail()
// 		}
// 	}()

// 	mang := managers.ComprobanteManager{
// 		// Ctx: ctx, // set this bar if mongo is deployed on replica set mode.
// 	}
// 	dataComprobante := models.Comprobante{
// 		TipoComprobante: nil
// 		Comprobante: "Nuevo",
// 		Descripcion: "nuevo comprobante",
// 		_id: "5e153d385cbbb2d76aafa1d3",
// 		Numero: 12221111
// 	}

// 	if err := mang.AddItem(&dataComprobante) err != nil {
// 		panic(err.Error())
// 	}

// }
