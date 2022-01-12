package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/udistrital/cuentas_contables_crud/compositors"
	"github.com/udistrital/cuentas_contables_crud/helpers"
	"github.com/udistrital/cuentas_contables_crud/models"
)

// ComprobanteController ...
type ComprobanteController struct {
	beego.Controller
	commonHelper          helpers.CommonHelper
	comprobanteCompositor compositors.ComprobanteCompositor
}

// var commonHelper = helpers.CommonHelper{}

// GetByUUID ...
// @Title Get
// @Description Obtiene los objetos por id
// @Param id path string true "UUID"
// @Success 200 {object} models.Comprobante
// @Failure 403 :objectId is empty
// @router /:id [get]
func (c *ComprobanteController) GetByUUID() {
	UUID := c.Ctx.Input.Param(":id")

	TipoComprobanteInfo, err := c.comprobanteCompositor.GetComprobanteByID(UUID)

	c.Data["json"] = c.commonHelper.DefaultResponse(200, err, TipoComprobanteInfo)

	c.ServeJSON()
}

// GetAll ...
// @Title Get
// @Description Obtiene todos los objetos
// @Success 200 {object} []models.Comprobante
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
// @Success 200 {string} "comprobante-added"
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
// @Param id   path string             true  "The objectid you want to update"
// @Param body body models.Comprobante true  "The new content"
// @Success 200 {string} "comprobante-updated"
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
// @Param id path string true "The objectid you want to delete"
// @Success 200 {string} "comprobante-deleted"
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
