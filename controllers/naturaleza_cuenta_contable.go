package controllers

import (
	"github.com/astaxie/beego"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/udistrital/cuentas_contables_crud/helpers"
	"github.com/udistrital/cuentas_contables_crud/managers"
	"github.com/udistrital/cuentas_contables_crud/models"
)

// NaturalezaCuentaContable ...
type NaturalezaCuentaContable struct {
	beego.Controller
	commonHelper helpers.CommonHelper
	crudManager  managers.CrudManager
}

// URLMapping ...
func (c *NaturalezaCuentaContable) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
}

// GetAll función para obtener todos los objetos
// @Title Get
// @Description get all objects
// @Success 200 {object} []models.NaturalezaCuentaContable
// @router / [get]
func (c *NaturalezaCuentaContable) GetAll() {
	filter := make(map[string]interface{})

	var responseData []*models.NaturalezaCuentaContable

	err := c.crudManager.GetAllDocuments(filter, -1, 0, models.NaturalezaCuentaContableCollection, func(curr *mongo.Cursor) {
		var row models.NaturalezaCuentaContable
		if err := curr.Decode(&row); err == nil {
			responseData = append(responseData, &row)
		}
	})

	c.Data["json"] = c.commonHelper.DefaultResponse(200, err, responseData)

	c.ServeJSON()
}
