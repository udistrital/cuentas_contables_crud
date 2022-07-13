package controllers

import (
	"github.com/astaxie/beego"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/udistrital/cuentas_contables_crud/helpers"
	"github.com/udistrital/cuentas_contables_crud/managers"
	"github.com/udistrital/cuentas_contables_crud/models"
)

// CentroCostos ...
type CentroCostos struct {
	beego.Controller
	commonHelper helpers.CommonHelper
	crudManager  managers.CrudManager
}

// URLMapping ...
func (c *CentroCostos) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
}

// GetAll ...
// @Title Get
// @Description Lista todos los centros de costos
// @Success 200 {object} models.CentroCostos
// @Failure 403 :objectId is empty
// @router / [get]
func (c *CentroCostos) GetAll() {
	filter := make(map[string]interface{})

	var responseData []*models.CentroCostos

	err := c.crudManager.GetAllDocuments(filter, -1, 0, models.CentroCostosCollection, func(curr *mongo.Cursor) {
		var row models.CentroCostos
		if err := curr.Decode(&row); err == nil {
			responseData = append(responseData, &row)
		}
	})

	c.Data["json"] = c.commonHelper.DefaultResponse(200, err, responseData)

	c.ServeJSON()
}
