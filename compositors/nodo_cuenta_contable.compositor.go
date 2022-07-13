package compositors

import (
	"context"

	"github.com/udistrital/cuentas_contables_crud/helpers"
	"github.com/udistrital/cuentas_contables_crud/managers"
	"github.com/udistrital/cuentas_contables_crud/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// NodoCuentaContableCompositor compositor for controll the data processing over NodoCuentaContable models
type NodoCuentaContableCompositor struct {
	crudManager   managers.CrudManager
	nodoCcManager managers.NodoCuentaContableManager
	nodoCcHelper  helpers.NodoCuentaContableHelper
}

// GetNodeByID Returns a *models.NodoCuentaContable by it's _id
func (c *NodoCuentaContableCompositor) GetNodeByID(ID string) (node *models.NodoCuentaContable, err error) {

	resul := &models.NodoCuentaContable{}

	id, _ := primitive.ObjectIDFromHex(ID)

	err = c.crudManager.GetDocumentByUUID(id, models.ArbolPlanMaestroCuentasContCollection, resul)

	return resul, err
}

// GetNodeByCode Returns a *models.NodoCuentaContable by it's code
func (c *NodoCuentaContableCompositor) GetNodeByCode(Codigo string) (node *models.NodoCuentaContable, err error) {

	resul := &models.NodoCuentaContable{}

	err = c.crudManager.GetDocumentByItem(Codigo, "codigo", models.ArbolPlanMaestroCuentasContCollection, resul)
	return resul, err
}

// GetNodeByNaturalezaCuentaContableC Returns a *models.NodoCuentaContable by it's naturaleza_id
func (c *NodoCuentaContableCompositor) GetNodeByNaturalezaCuentaContable(NaturalezaCuentaContable string, withNoActive ...bool) (rootNodes []*models.ArbolNbFormatNode, err error) {

	rootNodes, _, err = c.nodoCcManager.GetRootNodes(NaturalezaCuentaContable, withNoActive...)

	if err != nil {
		return
	}
	_, noRootNodes, err := c.nodoCcManager.GetNoRootNodes(NaturalezaCuentaContable, withNoActive...)
	if err != nil {
		return
	}
	c.nodoCcHelper.BuildTreeFromDataSource(rootNodes, noRootNodes)
	return
}

// GetCuentas Returns a []*models.ArkaCuentasContables
func (c *NodoCuentaContableCompositor) GetAll(filter map[string]interface{}, limit int64, offset int64) (nodesData []*models.ArkaCuentasContables, err error) {
	err = c.crudManager.GetAllDocuments(filter, limit, offset, models.ArbolPlanMaestroCuentasContCollection, func(curr *mongo.Cursor) {
		var node models.ArkaCuentasContables
		if err := curr.Decode(&node); err == nil {
			nodesData = append(nodesData, &node)
		}
	})
	return nodesData, err
}

// GetNodeArka Returns a *models.ArkaCuentasContables by it's naturaleza_id
func (c *NodoCuentaContableCompositor) GetNodeArka(NaturalezaCuentaContable string) (nodesData []*models.ArkaCuentasContables, err error) {

	filter := make(map[string]interface{})

	if NaturalezaCuentaContable != "" {
		filter["naturaleza_id"] = NaturalezaCuentaContable
	}
	err = c.crudManager.GetAllDocuments(filter, -1, 0, models.ArbolPlanMaestroCuentasContCollection, func(curr *mongo.Cursor) {
		var node models.ArkaCuentasContables
		if err := curr.Decode(&node); err == nil {
			nodesData = append(nodesData, &node)
		}
	})
	return nodesData, err

}

// AddNode Add new node to the tree
func (c *NodoCuentaContableCompositor) AddNode(nodeData *models.NodoCuentaContable) (err error) {

	err = c.crudManager.RunTransaction(func(ctx context.Context) error {
		ccmang := managers.NewNodoCuentaContableManager(nil)
		err = ccmang.AddNode(nodeData)
		return err
	})
	return
}

// BuildTree returns the tree data on the DB as a tree structure with it's hierarchy
func (c *NodoCuentaContableCompositor) BuildTree(withNoActive ...bool) (rootNodes []*models.ArbolNbFormatNode, err error) {
	rootNodes, _, err = c.nodoCcManager.GetRootNodes("", withNoActive...)

	if err != nil {
		return
	}
	_, noRootNodes, err := c.nodoCcManager.GetNoRootNodes("", withNoActive...)
	if err != nil {
		return
	}
	c.nodoCcHelper.BuildTreeFromDataSource(rootNodes, noRootNodes)
	return
}

// DeleteNode Add new node to the tree
func (c *NodoCuentaContableCompositor) DeleteNode(nodeData *models.NodoCuentaContable) (err error) {

	err = c.crudManager.RunTransaction(func(ctx context.Context) error {
		ccmang := managers.NewNodoCuentaContableManager(nil)
		err = ccmang.DeleteNodeByUUID(nodeData.ID)
		return err
	})
	return
}
