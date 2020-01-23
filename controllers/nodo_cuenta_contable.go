package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/udistrital/cuentas_contables_crud/compositors"
	"github.com/udistrital/cuentas_contables_crud/helpers"
	"github.com/udistrital/cuentas_contables_crud/managers"
	"github.com/udistrital/cuentas_contables_crud/models"
)

// NodoCuentaContableController ...
type NodoCuentaContableController struct {
	beego.Controller
	nodeCCCompositor compositors.NodoCuentaContableCompositor
	commonHelper     helpers.CommonHelper
	nodeCCManager    managers.NodoCuentaContableManager
	crudManager      managers.CrudManager
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
	if v, err := c.GetBool("fullTree"); v && err == nil {
		fullTree = v
	}
	treeData, err := c.nodeCCCompositor.BuildTree(fullTree)
	c.Data["json"] = c.commonHelper.DefaultResponse(200, err, treeData)

	c.ServeJSON()
}

// ChangeNodeState Método PUT de HTTP
// @Title PUT ChangeNodeState
// @Description Post models.NodoCuentaContable
// @Param	UUID		path 	string	true		"The key for object to update state"
// @Success 200 {int} models.NodoCuentaContable.Id
// @Failure 403 body is empty
// @router /change_node_state/:UUID [put]
func (c *NodoCuentaContableController) ChangeNodeState() {
	uuid := c.GetString(":UUID")
	/*
		   TODO: currently, this funtion will only change state
				 of target node, in future realises maybe it can
				 change full branch state.
	*/
	err := c.nodeCCManager.ChangeNodeState(uuid)
	message := ""
	if err == nil {
		message = "node-state-changed"
	}
	c.Data["json"] = c.commonHelper.DefaultResponse(200, err, message)
	c.ServeJSON()
}

// UpdateNode Método PUT de HTTP
// @Title PUT UpdateNode
// @Description Post models.NodoCuentaContable
// @Param	UUID		path 	string	true		"The key for object to update"
// @Success 200 {int} models.NodoCuentaContable.Id
// @Failure 403 body is empty
// @router /:UUID [put]
func (c *NodoCuentaContableController) UpdateNode() {
	uuid := c.GetString(":UUID")
	var requestBody models.NodoCuentaContable

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	message := ""

	if err == nil {
		requestBody.ID = uuid
		var resul interface{}
		err = c.crudManager.UpdateDocument(requestBody, uuid, models.ArbolPlanMaestroCuentasContCollection, &resul)
		if err == nil {
			message = "node-updated"
		}
	} else {
		message = "invalid-body"
	}

	c.Data["json"] = c.commonHelper.DefaultResponse(200, err, message)
	c.ServeJSON()
}
