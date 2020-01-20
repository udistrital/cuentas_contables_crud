package comprobantestestcases

import (
	"testing"

	"github.com/udistrital/cuentas_contables_crud/managers"
	"github.com/udistrital/cuentas_contables_crud/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	crudManager := managers.CrudManager{
		// Ctx: ctx, // set this bar if mongo is deployed on replica set mode.
	}
	if err := mang.AddItem(&dataComprobante); err != nil {
		panic(err.Error())
	}
	var comprobanteTest *models.Comprobante
	if err := crudManager.GetDocumentByItem(dataComprobante.Comprobante, "comprobante", models.ComprobanteCollection, &comprobanteTest); err != nil {
		panic(err.Error())
	}
	var updtDoc interface{}
	objectID, _ := primitive.ObjectIDFromHex(comprobanteTest.ID)
	if err := crudManager.DeleteDocumentByUUID(objectID, models.ComprobanteCollection, updtDoc); err != nil {
		panic(err.Error())
	}

}

func TestComprobanteFail(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Log("TestComprobanteFail Finalizado Correctamente (OK)")
		} else {
			t.Error("error: Comprobante doesn't create ")
			t.Fail()
		}
	}()

	mang := managers.ComprobanteManager{
		// Ctx: ctx, // set this bar if mongo is deployed on replica set mode.
	}
	dataComprobante := models.Comprobante{
		TipoComprobante: nil,
		Comprobante:     "Nuevo",
		Descripcion:     "nuevo comprobante",
		ID:              "5e153d385cbbb2d76aafa1d3",
		Numero:          12221111,
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

	dataComprobante = models.Comprobante{
		TipoComprobante:      &tipoComprobante,
		Comprobante:          "asdjs",
		Descripcion:          "askjdlas",
		Numero:               21312312,
		FechaCreacion:        "2020-01-08",
		Activo:               true,
		NumInicial:           213123,
		CuentaBanco:          "adasdas",
		TipoImpresion:        "Hojas normales",
		FormatoImpresion:     "Formato Adicional",
		NumCopias:            12312,
		Titulo:               "asdsad",
		NumeracionAutomatica: true,
	}

	if err := mang.UpdateItem(&dataComprobante, comprobanteTest.ID); err != nil {
		panic(err.Error())
	}
	var comprobanteTest *models.Comprobante
	if err := crudManager.GetDocumentByItem(dataComprobante.Comprobante, "comprobante", models.ComprobanteCollection, &comprobanteTest); err != nil {
		panic(err.Error())
	}
	var updtDoc interface{}
	objectID, _ := primitive.ObjectIDFromHex(comprobanteTest.ID)
	if err := crudManager.DeleteDocumentByUUID(objectID, models.ComprobanteCollection, updtDoc); err != nil {
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
	tipoComprobante := models.TipoComprobante{
		TipoDocumento: "A",
		Descripcion:   "Comprobante Put Egreso A",
	}

	dataComprobante := models.Comprobante{
		TipoComprobante:      &tipoComprobante,
		Comprobante:          "asdjs",
		Descripcion:          "askjdlas",
		Numero:               21312312,
		FechaCreacion:        "2020-01-08",
		Activo:               true,
		NumInicial:           213123,
		_id:                  "5e153d385cbbb2d76aafa1d3",
		CuentaBanco:          "adasdas",
		TipoImpresion:        "Hojas normales",
		FormatoImpresion:     "Formato Adicional",
		NumCopias:            12312,
		Titulo:               "asdsad",
		NumeracionAutomatica: true,
	}
	if err := mang.UpdateItem(&dataComprobante, comprobanteTest.ID); err != nil {
		var updtDoc interface{}
		objectID, _ := primitive.ObjectIDFromHex(comprobanteTest.ID)
		if err := crudManager.DeleteDocumentByUUID(objectID, models.ComprobanteCollection, updtDoc); err != nil {
			panic(err.Error())
		}
		panic(err.Error())
	}
}
