package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/udistrital/cuentas_contables_crud/compositors"
	_ "github.com/udistrital/cuentas_contables_crud/helpers"
	"github.com/udistrital/cuentas_contables_crud/models"
)

type TipoComprobanteController struct {
	beego.Controller
}

var tipoComprobanteCompositor = compositors.TipoComprobanteCompositor{}

// var commonHelper = helpers.CommonHelper{}

// GetByUUID función para obtener los objetos por id
// @Title Get
// @Description get object by id
// @Success 200 TipoComprobante models.TipoComprobante
// @Failure 403 :objectId is empty
// @router /:UUID [get]
func (c *TipoComprobanteController) GetByUUID() {
	UUID := c.Ctx.Input.Param(":UUID")

	TipoComprobanteInfo, err := tipoComprobanteCompositor.GetTipoComprobanteByID(UUID)

	c.Data["json"] = commonHelper.DefaultResponse(200, err, TipoComprobanteInfo)

	c.ServeJSON()
}

// GetAll función para obtener todos los objetos
// @Title Get
// @Description get all objects
// @Success 200 TipoComprobante models.TipoComprobante
// @Failure 403 :objectId is empty
// @router / [get]
func (c *TipoComprobanteController) GetAll() {

	TipoComprobanteInfo, err := tipoComprobanteCompositor.GetAllTipoComprobante()

	c.Data["json"] = commonHelper.DefaultResponse(200, err, TipoComprobanteInfo)

	c.ServeJSON()
}

// AddTipoComprobante Método Post de HTTP
// @Title Post models.TipoComprobante
// @Description Post models.TipoComprobante
// @Param	body		body 	models.TipoComprobante	true		"Body para la creacion de models.TipoComprobante"
// @Success 200 {int} models.TipoComprobante.Id
// @Failure 403 body is empty
// @router / [post]
func (c *TipoComprobanteController) AddTipoComprobante() {
	var requestBody models.TipoComprobante

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)

	message := ""

	if err == nil {
		err = tipoComprobanteCompositor.AddTipoComprobante(&requestBody)
	}

	if err == nil {
		message = "tipo-comprobante-added"
	}

	c.Data["json"] = commonHelper.DefaultResponse(200, err, message)

	c.ServeJSON()
}

// UpdateTipoComprobante Método Put de HTTP
// @Title Update models.TipoComprobante
// @Description Update models.TipoComprobante
// @Param	body		body 	models.TipoComprobante	true		"The objectid you want to update"
// @Success 200 {int} models.TipoComprobante.Id
// @Failure 403 body is empty
// @router /:id [put]
func (c *TipoComprobanteController) UpdateTipoComprobante() {
	objectID := c.Ctx.Input.Param(":id")
	var requestBody models.TipoComprobante

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)

	message := ""

	if err == nil {
		err = tipoComprobanteCompositor.UpdateTipoComprobante(&requestBody, objectID)
	}

	if err == nil {
		message = "tipo-comprobante-updated"
	}

	c.Data["json"] = commonHelper.DefaultResponse(200, err, message)

	c.ServeJSON()
}

// DeleteTipoComprobante Método Delete de HTTP
// @Title Delete models.TipoComprobante
// @Description Delete models.TipoComprobante
// @Param	body		body 	models.TipoComprobante	true		"The objectid you want to delete"
// @Success 200 {int} models.TipoComprobante.Id
// @Failure 403 body is empty
// @router /:id [delete]
func (c *TipoComprobanteController) DeleteTipoComprobante() {
	objectID := c.Ctx.Input.Param(":id")

	message := ""

	err := tipoComprobanteCompositor.DeleteTipoComprobante(objectID)

	if err == nil {
		message = "tipo-comprobante-deleted"
	}

	c.Data["json"] = commonHelper.DefaultResponse(200, err, message)

	c.ServeJSON()
}
