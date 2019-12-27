package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/udistrital/cuentas_contables_crud/compositors"
	"github.com/udistrital/cuentas_contables_crud/models"
)

type NodoCuentaContableController struct {
	beego.Controller
}

var nodeCCCompositor = compositors.NodoCuentaContableCompositor{}

// GetByUUID función para obtener todos los objetos
// @Title Get
// @Description get all objects
// @Success 200 NodoRubroApropiacion models.NodoCuentaContable
// @Failure 403 :objectId is empty
// @router /:UUID [get]
func (c *NodoCuentaContableController) GetByUUID() {
	UUID := c.GetString("UUID")

	nodeInfo, err := nodeCCCompositor.GetNodeByID(UUID, models.ArbolPlanMaestroCuentasContCollection)
	errorMessage := ""
	if err != nil {
		errorMessage = err.Error()
	}
	c.Data["json"] = map[string]interface{}{"data": nodeInfo, "error": errorMessage}
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
	errorMessage := ""
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err == nil {
		err = nodeCCCompositor.AddNode(&requestBody, models.ArbolPlanMaestroCuentasContCollection)
	}
	if err != nil {
		errorMessage = err.Error()
	}
	c.Data["json"] = map[string]interface{}{"data": requestBody, "error": errorMessage}
	c.ServeJSON()
}
