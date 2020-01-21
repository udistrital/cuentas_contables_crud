package comprobantestestcases

import (
	"testing"
	"time"

	"github.com/udistrital/cuentas_contables_crud/managers"
	"github.com/udistrital/cuentas_contables_crud/models"
)

func TestComprobanteSuccess(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Error("error: ", r)
			t.Fail()
		} else {
			t.Log("TestTrComprobante Finalizado Correctamente (OK)")
		}
	}()

	tipoComprobante := models.TipoComprobante{
		TipoDocumento: "G",
		Descripcion:   "Comprobante Put Egreso",
	}

	dataComprobante := models.Comprobante{
		TipoComprobante: &tipoComprobante,
		Comprobante:     "testcomprobante",
		Descripcion:     "testcomprobante",
		Numero:          12221111,
	}

	mang := managers.ComprobanteManager{
		// Ctx: ctx, // set this bar if mongo is deployed on replica set mode.
	}
	if err := mang.AddItem(&dataComprobante); err != nil {
		panic(err.Error())
	}
}

func TestComprobanteFail(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Log("TestComprobanteFail Finalizado Correctamente (OK)")
		} else {
			t.Error("error: Comprobante created ")
			t.Fail()
		}
	}()

	mang := managers.ComprobanteManager{
		// Ctx: ctx, // set this bar if mongo is deployed on replica set mode.
	}
	dataComprobante := models.Comprobante{

		Comprobante: "Nuevo",
		Descripcion: "nuevo comprobante",
		ID:          "111",
		Numero:      12221111,
	}

	if err := mang.AddItem(&dataComprobante); err != nil {
		panic(err.Error())
	}

}

func TestComprobanteParamSuccess(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Error("error: ", r)
			t.Fail()
		} else {
			t.Log("TestTrComprobanteParams Finalizado Correctamente (OK)")
		}
	}()
	tipoComprobante := models.TipoComprobante{
		TipoDocumento: "C",
		Descripcion:   "Comprobante Put Egreso C",
	}

	dataComprobante := models.Comprobante{
		TipoComprobante: &tipoComprobante,
		Comprobante:     "test",
		Descripcion:     "test",
		Numero:          99999,
	}
	mang := managers.ComprobanteManager{
		// Ctx: ctx, // set this bar if mongo is deployed on replica set mode.
	}
	if err := mang.AddItem(&dataComprobante); err != nil {
		panic(err.Error())
	}
	crudManager := managers.CrudManager{}
	var comprobanteTest *models.Comprobante
	if err := crudManager.GetDocumentByItem(dataComprobante.Comprobante, "comprobante", models.ComprobanteCollection, &comprobanteTest); err != nil {
		panic(err.Error())
	}
	tipoComprobante = models.TipoComprobante{
		TipoDocumento: "A",
		Descripcion:   "Comprobante Put Egreso A",
	}
	general := models.General{
		FechaCreacion: time.Now().Format("2006-01-02T15:04:05"),
		Activo:        true,
	}

	parametros := models.Parametros{
		NumInicial:           213123,
		CuentaBanco:          "adasdas",
		TipoImpresion:        "Hojas normales",
		FormatoImpresion:     "Formato Adicional",
		NumCopias:            12312,
		Titulo:               "asdsad",
		NumeracionAutomatica: true,
	}

	dataComprobante = models.Comprobante{
		TipoComprobante: &tipoComprobante,
		Comprobante:     "asdjs",
		Descripcion:     "askjdlas",
		Numero:          21312312,
		General:         &general,
		Parametros:      &parametros,
	}

	if err := mang.UpdateItem(&dataComprobante, comprobanteTest.ID); err != nil {
		panic(err.Error())
	}

}
func TestComprobanteParamFail(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Log("TestComprobanteParamFail Finalizado Correctamente (OK)")
		} else {
			t.Error("error: Comprobante doesn't update ")
			t.Fail()
		}
	}()
	tipoComprobante := models.TipoComprobante{
		TipoDocumento: "C",
		Descripcion:   "Comprobante Put Egreso C",
	}

	dataComprobante := models.Comprobante{
		TipoComprobante: &tipoComprobante,
		Comprobante:     "test",
		Descripcion:     "test",
		Numero:          99999,
	}
	mang := managers.ComprobanteManager{
		// Ctx: ctx, // set this bar if mongo is deployed on replica set mode.
	}
	if err := mang.AddItem(&dataComprobante); err != nil {
		panic(err.Error())
	}
	crudManager := managers.CrudManager{}
	var comprobanteTest *models.Comprobante
	if err := crudManager.GetDocumentByItem(dataComprobante.Comprobante, "comprobante", models.ComprobanteCollection, &comprobanteTest); err != nil {
		panic(err.Error())
	}
	tipoComprobante = models.TipoComprobante{
		TipoDocumento: "A",
		Descripcion:   "Comprobante Put Egreso A",
	}
	general := models.General{
		FechaCreacion: time.Now().Format("2006-01-02T15:04:05"),
		Activo:        true,
	}

	parametros := models.Parametros{
		NumInicial:           213123,
		CuentaBanco:          "adasdas",
		TipoImpresion:        "Hojas normales",
		FormatoImpresion:     "Formato Adicional",
		NumCopias:            12312,
		Titulo:               "asdsad",
		NumeracionAutomatica: true,
	}
	dataComprobante = models.Comprobante{
		TipoComprobante: &tipoComprobante,
		Comprobante:     "asdjs",
		Descripcion:     "askjdlas",
		Numero:          21312312,
		General:         &general,
		Parametros:      &parametros,
	}
	if err := mang.UpdateItem(&dataComprobante, "A"); err != nil {
		panic(err.Error())
	}
}
