package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/udistrital/cuentas_contables_crud/compositors"
	"github.com/udistrital/cuentas_contables_crud/helpers"
	_ "github.com/udistrital/cuentas_contables_crud/helpers"
	"github.com/udistrital/cuentas_contables_crud/models"
)

type ComprobanteController struct {
	beego.Controller
	commonHelper          helpers.CommonHelper
	comprobanteCompositor compositors.ComprobanteCompositor
}

// var commonHelper = helpers.CommonHelper{}

// GetByUUID función para obtener los objetos por id
// @Title Get
// @Description get object por id
// @Success 200 Comprobante models.Comprobante
// @Failure 403 :objectId is empty
// @router /:UUID [get]
func (c *ComprobanteController) GetByUUID() {
	UUID := c.Ctx.Input.Param(":UUID")

	TipoComprobanteInfo, err := c.comprobanteCompositor.GetComprobanteByID(UUID)

	c.Data["json"] = c.commonHelper.DefaultResponse(200, err, TipoComprobanteInfo)

	c.ServeJSON()
}

// GetAll función para obtener todos los objetos
// @Title Get
// @Description get all objects
// @Success 200 Comprobante models.Comprobante
// @Failure 403 :objectId is empty
// @router / [get]
func (c *ComprobanteController) GetAll() {

	TipoComprobanteInfo, err := c.comprobanteCompositor.GetAllComprobante()

	c.Data["json"] = c.commonHelper.DefaultResponse(200, err, TipoComprobanteInfo)

	c.ServeJSON()
}

// AddComprobante Método Post de HTTP
// @Title Post models.Comprobante
// @Description Post models.Comprobante
// @Param	body		body 	models.Comprobante	true		"Body para la creacion de models.Comprobante"
// @Success 200 {int} models.Comprobante.Id
// @Failure 403 body is empty
// @router / [post]
func (c *ComprobanteController) AddComprobante() {
	var requestBody models.Comprobante

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)

	message := ""

	if err == nil {
		err = c.comprobanteCompositor.AddComprobante(&requestBody)
	}

	if err == nil {
		message = "comprobante-added"
	}

	c.Data["json"] = c.commonHelper.DefaultResponse(200, err, message)

	c.ServeJSON()
}

// UpdateComprobante Método Put de HTTP
// @Title Update models.Comprobante
// @Description Update models.Comprobante
// @Param	body		body 	models.Comprobante	true		"The objectid you want to update"
// @Success 200 {int} models.Comprobante.Id
// @Failure 403 body is empty
// @router /:id [put]
func (c *ComprobanteController) UpdateComprobante() {
	objectID := c.Ctx.Input.Param(":id")
	var requestBody models.Comprobante

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)

	message := ""

	if err == nil {
		err = c.comprobanteCompositor.UpdateComprobante(&requestBody, objectID)
	}

	if err == nil {
		message = "comprobante-updated"
	}

	c.Data["json"] = c.commonHelper.DefaultResponse(200, err, message)

	c.ServeJSON()
}

// DeleteComprobante Método Delete de HTTP
// @Title Delete models.Comprobante
// @Description Delete models.Comprobante
// @Param	body		body 	models.Comprobante	true		"The objectid you want to delete"
// @Success 200 {int} models.Comprobante.Id
// @Failure 403 body is empty
// @router /:id [delete]
func (c *ComprobanteController) DeleteComprobante() {
	objectID := c.Ctx.Input.Param(":id")

	message := ""

	err := c.comprobanteCompositor.DeleteComprobante(objectID)

	if err == nil {
		message = "comprobante-deleted"
	}

	c.Data["json"] = c.commonHelper.DefaultResponse(200, err, message)

	c.ServeJSON()
}
