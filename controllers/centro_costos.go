package controllers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/cuentas_contables_crud/helpers"
	"github.com/udistrital/cuentas_contables_crud/managers"
	"github.com/udistrital/cuentas_contables_crud/models"
	"go.mongodb.org/mongo-driver/mongo"
)

// CentroCostos ...
type CentroCostos struct {
	beego.Controller
	commonHelper helpers.CommonHelper
	crudManager  managers.CrudManager
}

// GetAll ...
// @Title Get
// @Description Lista todos los centros de costos
// @Success 200 {object} models.TipoComprobante
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
