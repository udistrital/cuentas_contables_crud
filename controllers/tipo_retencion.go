package controllers

import (
	"github.com/astaxie/beego"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/udistrital/cuentas_contables_crud/helpers"
	"github.com/udistrital/cuentas_contables_crud/managers"
	"github.com/udistrital/cuentas_contables_crud/models"
)

// TipoRetencion ...
type TipoRetencionController struct {
	beego.Controller
	commonHelper helpers.CommonHelper
	crudManager  managers.CrudManager
}

// GetAll función para obtener todos los objetos
// @Title Get
// @Description get all objects
// @Success 200 {object} []models.TipoRetencion
// @router / [get]
func (c *TipoRetencionController) GetAll() {
	filter := make(map[string]interface{})

	var responseData []*models.TipoRetencion

	err := c.crudManager.GetAllDocuments(filter, -1, 0, models.TipoRetencionCollection, func(curr *mongo.Cursor) {
		var row models.TipoRetencion
		if err := curr.Decode(&row); err == nil {
			responseData = append(responseData, &row)
		}
	})

	c.Data["json"] = c.commonHelper.DefaultResponse(200, err, responseData)

	c.ServeJSON()
}
