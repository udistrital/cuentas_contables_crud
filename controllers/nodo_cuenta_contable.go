package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/cuentas_contables_crud/compositors"
	"github.com/udistrital/cuentas_contables_crud/helpers"
	"github.com/udistrital/cuentas_contables_crud/managers"
	"github.com/udistrital/cuentas_contables_crud/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// GetCuentasUsablesByNaturaleza obtiene las cuentas de máximo nivel (sin hijos) segun su naturaleza
// @Title GetCuentas
// @Description obtiene las cuentas de máximo nivel (sin hijos) segun su naturaleza
// @Param	NaturalezaCuentaContable		path 	string	true	"NaturalezaCuentaContable para el filtro por tipo de cuenta contable(credito/debito)"
// @Param	withInactives	query	bool	false	"With inactives nodes. False is default"
// @Success 200  models.ArkaCuentasContables
// @Failure 403 :objectId is empty
// @router /getCuentas/:NaturalezaCuentaContable [get]
func (c *NodoCuentaContableController) GetCuentasUsablesByNaturaleza() {
	NaturalezaCuentaContable := c.GetString(":NaturalezaCuentaContable")
	withInactives := false
	if v, err := c.GetBool("withInactives"); v && err == nil {
		withInactives = v
	}
	filter := make(map[string]interface{})
	if NaturalezaCuentaContable != "" {
		filter["naturaleza_id"] = NaturalezaCuentaContable
	}
	if !withInactives {
		filter["activo"] = true
	}
	filter["$or"] = []bson.M{{"hijos": nil}, {"hijos": []bson.M{}}}
	nodeInfo, err := c.nodeCCCompositor.GetAll(filter, -1, 0)
	c.Data["json"] = c.commonHelper.DefaultResponse(200, err, nodeInfo)
	c.ServeJSON()
}

// GetByNaturalezaArka función para obtener Los objetos segun naturaleza de cuenta contable para consumir en arka
// @Title Get
// @Description get all objects based on naturaleza cuenta contable for arka client
// @Param	NaturalezaCuentaContable		path 	string	true	"NaturalezaCuentaContable para el filtro por tipo de cuenta contable(credito/debito)"
// @Success 200  models.ArkaCuentasContables
// @Failure 403 :objectId is empty
// @router /getNodosCuentasArka/:NaturalezaCuentaContable [get]
func (c *NodoCuentaContableController) GetByNaturalezaArka() {
	NaturalezaCuentaContable := c.GetString(":NaturalezaCuentaContable")

	nodeInfo, err := c.nodeCCCompositor.GetNodeArka(NaturalezaCuentaContable)

	c.Data["json"] = c.commonHelper.DefaultResponse(200, err, nodeInfo)

	c.ServeJSON()

}

// GetByNaturalezaCuentaContable función para obtener Los objetos segun naturaleza de cuenta contable
// @Title Get
// @Description get all objects based on naturaleza cuenta contable
// @Param	NaturalezaCuentaContable		path 	string	true	"NaturalezaCuentaContable para el filtro por tipo de cuenta contable(credito/debito)"
// @Success 200 NodoRubroApropiacion models.NodoCuentaContable
// @Failure 403 :objectId is empty
// @router /cuentas/:NaturalezaCuentaContable [get]
func (c *NodoCuentaContableController) GetByNaturalezaCuentaContable() {
	NaturalezaCuentaContable := c.GetString(":NaturalezaCuentaContable")
	fullTree := false
	treeData, err := c.nodeCCCompositor.GetNodeByNaturalezaCuentaContable(NaturalezaCuentaContable, fullTree)
	c.Data["json"] = c.commonHelper.DefaultResponse(200, err, treeData)
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

// GetAll ...
// @Title Get All
// @Description Obtiene cuentas contables
// @Param	query	query	string	false	"Filter. e.g. {"naturaleza_id":"credito"}"
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Param	withInactives	query	bool	false	"With inactives nodes. False is default"
// @Success 200 {object} models.ArkaCuentasContables
// @Failure 404 not found resource
// @router /cuentas [get]
func (c *NodoCuentaContableController) GetAll() {
	var query bson.M = nil
	var limit int64 = -1
	var offset int64 = 0
	withInactives := false
	if v, err := c.GetBool("withInactives"); v && err == nil {
		withInactives = v
	}
	if v := c.GetString("query"); v != "" {
		err := json.Unmarshal([]byte(v), &query)
		if err != nil {
			logs.Error("json. Unmarshal() ERROR:", err)
		} else if _, exist := query["activo"]; !exist && !withInactives {
			query["activo"] = true
		}
	} else if !withInactives {
		query = bson.M{"activo": true}
	}
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	l, err := c.nodeCCCompositor.GetAll(query, limit, offset)
	if err != nil {
		logs.Error(err)
		c.Data["system"] = err
		c.Abort("404")
	} else {
		c.Data["json"] = c.commonHelper.DefaultResponse(200, err, l)
	}
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
		requestBody.ID, _ = primitive.ObjectIDFromHex(uuid)
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
