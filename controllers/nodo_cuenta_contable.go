package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/udistrital/cuentas_contables_crud/compositors"
	"github.com/udistrital/cuentas_contables_crud/helpers"
	"github.com/udistrital/cuentas_contables_crud/models"
)

type NodoCuentaContableController struct {
	beego.Controller
	nodeCCCompositor compositors.NodoCuentaContableCompositor
	commonHelper     helpers.CommonHelper
}

// GetByUUID función para obtener todos los objetos
// @Title Get
// @Description get all objects
// @Success 200 NodoRubroApropiacion models.NodoCuentaContable
// @Failure 403 :objectId is empty
// @router /:UUID [get]
func (c *NodoCuentaContableController) GetByUUID() {
	UUID := c.GetString(":UUID")

	nodeInfo, err := c.nodeCCCompositor.GetNodeByID(UUID)

	c.Data["json"] = c.commonHelper.DefaultResponse(200, err, nodeInfo)

	c.ServeJSON()
}

// AddNode Método Post de HTTP
// @Title Post models.NodoCuentaContable
// @Description Post models.NodoCuentaContable
// @Param	body		body 	models.NodoCuentaContable	true		"Body para la creacion de models.NodoCuentaContable"
// @Success 200 {int} models.NodoCuentaContable.Id
// @Failure 403 body is empty
// @router / [post]
func (c *NodoCuentaContableController) AddNode() {
	var requestBody models.NodoCuentaContable

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)

	message := ""

	if err == nil {
		err = c.nodeCCCompositor.AddNode(&requestBody)
	}

	if err == nil {
		message = "node-added"
	}

	c.Data["json"] = c.commonHelper.DefaultResponse(200, err, message)

	c.ServeJSON()
}

// GetTree función para obtener todos los objetos
// @Title GetTree
// @Description get all objects
// @Param	query	fullTree	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Success 200 NodoRubroApropiacion []models.NodoCuentaContable
// @Failure 403 :objectId is empty
// @router / [get]
func (c *NodoCuentaContableController) GetTree() {
	fullTree := false
	if v, _ := c.GetBool("fullTree"); v {
		fullTree = v
	}
	treeData, err := c.nodeCCCompositor.BuildTree(fullTree)
	c.Data["json"] = c.commonHelper.DefaultResponse(200, err, treeData)

	c.ServeJSON()
}
