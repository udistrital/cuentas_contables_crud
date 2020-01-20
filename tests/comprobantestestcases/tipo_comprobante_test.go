package comprobantestestcases

import (
	"testing"

	"github.com/udistrital/cuentas_contables_crud/managers"
	"github.com/udistrital/cuentas_contables_crud/models"
)

func TestTipoComprobanteSuccess(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Error("error: ", r)
			t.Fail()
		} else {
			t.Log("TestTrTipoComprobante Finalizado Correctamente (OK)")
		}
	}()

	dataTipoComprobante := models.TipoComprobante{
		TipoDocumento: "X",
		Descripcion:   "Comprobante X ejemplo",
	}

	mang := managers.TipoComprobanteManager{
		// Ctx: ctx, // set this bar if mongo is deployed on replica set mode.
	}
	if err := mang.AddItem(&dataTipoComprobante); err != nil {
		panic(err.Error())
	}

}

func TestTipoComprobanteFail(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Log("TestTipoComprobanteFail Finalizado Correctamente (OK)")
		} else {
			t.Error("error: TipoComprobante doesn't create ")
			t.Fail()
		}
	}()

	dataTipoComprobante := models.TipoComprobante{
		TipoDocumento: "Y",
		Descripcion:   "Comprobante Y ejemplo",
	}

	mang := managers.TipoComprobanteManager{
		// Ctx: ctx, // set this bar if mongo is deployed on replica set mode.
	}
	if err := mang.AddItem(&dataTipoComprobante); err != nil {
		panic(err.Error())
	}
	dataTipoComprobante2 := models.TipoComprobante{
		TipoDocumento: "Y",
		Descripcion:   "Comprobante Y ejemplo",
	}

	if err := mang.AddItem(&dataTipoComprobante2); err != nil {
		panic(err.Error())
	}

}
