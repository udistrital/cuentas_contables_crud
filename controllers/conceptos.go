package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/logs"
	"github.com/udistrital/cuentas_contables_crud/compositors"
	"github.com/udistrital/cuentas_contables_crud/helpers"
	"github.com/udistrital/cuentas_contables_crud/managers"
	"github.com/udistrital/cuentas_contables_crud/models"
	"github.com/udistrital/utils_oas/errorctrl"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ConceptosController ...
type ConceptosController struct {
	beego.Controller
	nodeConcCompositor compositors.ConceptosCompositor
	commonHelper       helpers.CommonHelper
	nodeConcCManager   managers.ConceptosManager
	crudManager        managers.CrudManager
}

// GetTree ...
// @Title GetTree
// @Description funcion para obtener todos los objetos
// @Success 200 {object} []models.ArbolConceptosFormatNode
// @Failure 400 :objectId is empty
// @router / [get]
func (c *ConceptosController) GetTree() {
	defer errorctrl.ErrorControlController(c.Controller, "ConceptosController")
	if treeData, err := c.nodeConcCompositor.BuildTreeConceptos(); err == nil {
		c.Data["json"] = c.commonHelper.DefaultResponse(200, err, treeData)
	} else {
		panic(errorctrl.Error("GetTree", err, "500"))
	}
	c.ServeJSON()
}

// GetByCodigo ...
// @Title GetByCodigo
// @Description get conceptos por codigo
// @Param	Codigo		path 	string	true		"Codigo del concepto"
// @Success 200 {object} models.Conceptos
// @Failure 400 :object is empty
// @router /:Codigo [get]
func (c *ConceptosController) GetByCodigo() {
	defer errorctrl.ErrorControlController(c.Controller, "ConceptosController")

	Codigo := c.GetString(":Codigo")
	if nodeInfo, err := c.nodeConcCompositor.GetNodeByCodigo(Codigo); err == nil {
		c.Data["json"] = c.commonHelper.DefaultResponse(200, err, nodeInfo)
	} else {
		panic(errorctrl.Error("GetByCodigo", err, "404"))
	}

	c.ServeJSON()

}

// AddNode ..
// @Title AddNode
// @Description create conceptos
// @Param	body		body 	models.Conceptos	true		"Body para la creacion de conceptos"
// @Success 201 {object} models.Conceptos
// @Failure 400 body is empty
// @router / [post]
func (c *ConceptosController) AddNode() {
	defer errorctrl.ErrorControlController(c.Controller, "ConceptosController")
	var requestBody models.Conceptos

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody); err == nil {
		if err = c.nodeConcCompositor.AddNodeConceptos(&requestBody); err == nil {
			c.Data["json"] = c.commonHelper.DefaultResponse(201, err, requestBody)
		} else {
			panic(errorctrl.Error("AddNode", err, "500"))
		}
	} else {
		panic(errorctrl.Error("AddNode", err, "500"))
	}
	c.ServeJSON()
}

// UpdateNode ...
// @Title PUT UpdateNode
// @Description Put conceptos
// @Param	UUID		path 	string	true		"The key for object to update"
// @Param	body		body 	models.Conceptos	true		"Body para la actualizacion de models.Conceptos"
// @Success 200 {object} models.Conceptos
// @Failure 400 body is empty
// @router /:UUID [put]
func (c *ConceptosController) UpdateNode() {
	defer errorctrl.ErrorControlController(c.Controller, "ConceptosController")
	uuid := c.GetString(":UUID")
	var requestBody models.Conceptos

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody); err != nil {
		panic(errorctrl.Error("UpdateNode", err, "400"))
	}
	var err error
	if requestBody.ID, err = primitive.ObjectIDFromHex(uuid); err != nil {
		panic(errorctrl.Error("UpdateNode", err, "400"))
	}
	var resul interface{}
	if err := c.crudManager.UpdateDocument(requestBody, requestBody.ID, models.ArbolConceptosCollection, &resul); err == nil {
		c.Data["json"] = c.commonHelper.DefaultResponse(200, err, requestBody)
	} else {
		panic(errorctrl.Error("UpdateNode", err, "500"))
	}
	c.ServeJSON()
}
