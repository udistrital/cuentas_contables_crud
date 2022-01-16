package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/logs"
	"github.com/udistrital/cuentas_contables_crud/compositors"
	"github.com/udistrital/cuentas_contables_crud/helpers"
	"github.com/udistrital/cuentas_contables_crud/managers"
	"github.com/udistrital/cuentas_contables_crud/models"

)

// ConceptosController ...
type ConceptosController struct {
	beego.Controller
	nodeConcCompositor compositors.ConceptosCompositor
	commonHelper     helpers.CommonHelper
	nodeConcCManager    managers.ConceptosManager
	crudManager      managers.CrudManager
}

// GetTree función para obtener todos los objetos
// @Title GetTree
// @Description get all objects
// @Param	query	fullTree	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Success 200 Conceptos []models.Conceptos
// @Failure 403 :objectId is empty
// @router / [get]
func (c *ConceptosController) GetTree() {
	fullTree := false
	if v, err := c.GetBool("fullTree"); v && err == nil {
		fullTree = v
	}
	treeData, err := c.nodeConcCompositor.BuildTreeConceptos(fullTree)
	c.Data["json"] = c.commonHelper.DefaultResponse(200, err, treeData)
	c.ServeJSON()
}

// GetByCodigo función para obtener todos los objetos
// @Title Get
// @Description get all objects
// @Success 200 Conceptos models.Conceptos
// @Failure 403 :objectId is empty
// @router /:Codigo [get]
func (c *ConceptosController) GetByCodigo() {
	Codigo := c.GetString(":Codigo")

	nodeInfo, err := c.nodeConcCompositor.GetNodeByCodigo(Codigo)

	c.Data["json"] = c.commonHelper.DefaultResponse(200, err, nodeInfo)

	c.ServeJSON()

}

// AddNode Método Post de HTTP
// @Title Post models.Conceptos
// @Description Post models.Conceptos
// @Param	body		body 	models.Conceptos	true		"Body para la creacion de models.Conceptos"
// @Success 200 {int} models.Conceptos.Id
// @Failure 403 body is empty
// @router / [post]
func (c *ConceptosController) AddNode() {
	var requestBody models.Conceptos

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)

	message := ""

	if err == nil {
		err = c.nodeConcCompositor.AddNodeConceptos(&requestBody)
	}

	if err == nil {
		message = "node-added"
	}

	c.Data["json"] = c.commonHelper.DefaultResponse(200, err, message)

	c.ServeJSON()
}
