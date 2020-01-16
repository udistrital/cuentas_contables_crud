package compositors

import (
	"context"

	"github.com/udistrital/cuentas_contables_crud/helpers"
	"github.com/udistrital/cuentas_contables_crud/managers"
	"github.com/udistrital/cuentas_contables_crud/models"
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

	err = c.crudManager.GetDocumentByUUID(ID, models.ArbolPlanMaestroCuentasContCollection, resul)

	return resul, err
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
	rootNodes, _, err = c.nodoCcManager.GetRootNodes(withNoActive...)

	if err != nil {
		return
	}
	_, noRootNodes, err := c.nodoCcManager.GetNoRootNodes(withNoActive...)
	if err != nil {
		return
	}
	c.nodoCcHelper.BuildTreeFromDataSource(rootNodes, noRootNodes)
	return
}
