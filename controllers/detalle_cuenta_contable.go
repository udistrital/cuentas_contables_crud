package controllers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/cuentas_contables_crud/helpers"
	"github.com/udistrital/cuentas_contables_crud/managers"
	"github.com/udistrital/cuentas_contables_crud/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type DetalleCuentaContable struct {
	beego.Controller
	commonHelper helpers.CommonHelper
	crudManager  managers.CrudManager
}

// GetAll función para obtener todos los objetos
// @Title Get
// @Description get all objects
// @Success 200 TipoComprobante models.TipoComprobante
// @Failure 403 :objectId is empty
// @router / [get]
func (c *DetalleCuentaContable) GetAll() {
	filter := make(map[string]interface{})

	var responseData []*models.DetalleCuentaContable

	err := c.crudManager.GetAllDocuments(filter, -1, 0, models.DetalleCuentaContableCollection, func(curr *mongo.Cursor) {
		var row models.DetalleCuentaContable
		if err := curr.Decode(&row); err == nil {
			responseData = append(responseData, &row)
		}
	})

	c.Data["json"] = c.commonHelper.DefaultResponse(200, err, responseData)

	c.ServeJSON()
}
