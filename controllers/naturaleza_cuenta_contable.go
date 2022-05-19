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
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} []models.NaturalezaCuentaContable
// @router / [get]
func (c *NaturalezaCuentaContable) GetAll() {
	filter := make(map[string]interface{})

	var responseData []*models.NaturalezaCuentaContable

	var limit int64 = -1
	var offset int64

	// limit: -1 (default is -1)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}

	err := c.crudManager.GetAllDocuments(filter, limit, offset, models.NaturalezaCuentaContableCollection, func(curr *mongo.Cursor) {
		var row models.NaturalezaCuentaContable
		if err := curr.Decode(&row); err == nil {
			responseData = append(responseData, &row)
		}
	})

	c.Data["json"] = c.commonHelper.DefaultResponse(200, err, responseData)

	c.ServeJSON()
}
