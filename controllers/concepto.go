package controllers

import (
	"encoding/json"
	"errors"

	"github.com/astaxie/beego"
	"github.com/udistrital/cuentas_contables_crud/helpers"
	"github.com/udistrital/cuentas_contables_crud/managers"
	"github.com/udistrital/cuentas_contables_crud/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ConceptoController struct {
	beego.Controller
	commonHelper helpers.CommonHelper
	crudManager  managers.CrudManager
}

// Add Método Post de HTTP
// @Title Post models.Concepto
// @Description Post models.Concepto
// @Param	body		body 	models.Concepto	true		"Body para la creacion de models.Concepto"
// @Success 200 {int} models.Concepto.Id
// @Failure 403 body is empty
// @router / [post]
func (c *ConceptoController) Add() {
	var requestBody models.Concepto

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)

	uuid := ""

	if err == nil {
		var checkConceptIntf interface{}
		if e := c.crudManager.GetDocumentByItem(requestBody.Nombre, "nombre", models.ConceptoCollection, &checkConceptIntf); e != nil {
			if e.Error() == "document-no-found-by-item" {
				uuid, err = c.crudManager.AddDocument(requestBody, models.ConceptoCollection)
			} else {
				err = e
			}
		} else {
			err = errors.New("concept-already-exist")
		}
	}

	if err == nil {
		requestBody.ID = uuid
	}

	c.Data["json"] = c.commonHelper.DefaultResponse(200, err, requestBody)

	c.ServeJSON()
}

// GetOne función para obtener todos los objetos
// @Title Get
// @Description get all objects
// @Success 200 TipoComprobante models.Concepto
// @Failure 403 :objectId is empty
// @router /:UUID [get]
func (c *ConceptoController) GetOne() {
	UUID := c.GetString(":UUID")
	objectID, _ := primitive.ObjectIDFromHex(UUID)
	var responseData models.Concepto

	err := c.crudManager.GetDocumentByUUID(objectID, models.ConceptoCollection, &responseData)

	c.Data["json"] = c.commonHelper.DefaultResponse(200, err, responseData)

	c.ServeJSON()
}

// GetAll función para obtener todos los objetos
// @Title Get
// @Description get all objects
// @Success 200 TipoComprobante models.TipoComprobante
// @Failure 403 :objectId is empty
// @router / [get]
func (c *ConceptoController) GetAll() {
	filter := make(map[string]interface{})

	var responseData []*models.Concepto

	err := c.crudManager.GetAllDocuments(filter, -1, 0, models.ConceptoCollection, func(curr *mongo.Cursor) {
		var row models.Concepto
		if err := curr.Decode(&row); err == nil {
			responseData = append(responseData, &row)
		}
	})

	c.Data["json"] = c.commonHelper.DefaultResponse(200, err, responseData)

	c.ServeJSON()
}

// UpdateNode Método PUT de HTTP
// @Title PUT UpdateNode
// @Description Post models.Concepto
// @Param	UUID		path 	string	true		"The key for object to update"
// @Success 200 {int} models.Concepto.ID
// @Failure 403 body is empty
// @router /:UUID [put]
func (c *ConceptoController) UpdateNode() {
	uuid := c.GetString(":UUID")
	var requestBody models.Concepto

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	message := ""

	if err == nil {
		var resul interface{}
		objectID, _ := primitive.ObjectIDFromHex(uuid)
		err = c.crudManager.UpdateDocument(requestBody, objectID, models.ConceptoCollection, &resul)
		if err == nil {
			message = "concept-updated"
		}
	} else {
		message = "invalid-body"
	}

	c.Data["json"] = c.commonHelper.DefaultResponse(200, err, message)
	c.ServeJSON()
}

// Delete Método PUT de HTTP
// @Title Delete Delete
// @Description Post models.Concepto
// @Param	UUID		path 	string	true		"The key for object to update"
// @Success 200 {int} models.Concepto.ID
// @Failure 403 body is empty
// @router /:UUID [delete]
func (c *ConceptoController) Delete() {
	uuid := c.GetString(":UUID")

	message := ""

	var resul interface{}
	objectID, _ := primitive.ObjectIDFromHex(uuid)

	err := c.crudManager.DeleteDocumentByUUID(objectID, models.ConceptoCollection, &resul)
	if err == nil {
		message = "concept-deleted"
	}

	c.Data["json"] = c.commonHelper.DefaultResponse(200, err, message)
	c.ServeJSON()
}
