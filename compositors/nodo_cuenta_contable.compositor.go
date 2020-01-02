package compositors

import (
	"context"

	"github.com/udistrital/cuentas_contables_crud/helpers"
	"github.com/udistrital/cuentas_contables_crud/managers"
	"github.com/udistrital/cuentas_contables_crud/models"
)

// NodoCuentaContableCompositor compositor for controll the data processing over NodoCuentaContable models
type NodoCuentaContableCompositor struct{}

var crudManager = managers.CrudManager{}
var nodoCcManager = managers.NodoCuentaContableManager{}
var nodoCcHelper = helpers.NodoCuentaContableHelper{}

// GetNodeByID Returns a *models.NodoCuentaContable by it's _id
func (c *NodoCuentaContableCompositor) GetNodeByID(ID string) (node *models.NodoCuentaContable, err error) {

	resul := &models.NodoCuentaContable{}

	err = crudManager.GetDocumentByUUID(ID, models.ArbolPlanMaestroCuentasContCollection, resul)

	return resul, err
}

// AddNode Add new node to the tree
func (c *NodoCuentaContableCompositor) AddNode(nodeData *models.NodoCuentaContable) (err error) {

	err = crudManager.RunTransaction(func(ctx context.Context) error {
		ccmang := managers.NewNodoCuentaContableManager(nil)
		err = ccmang.AddNode(nodeData)
		return err
	})
	return
}

// BuildTree returns the tree data on the DB as a tree structure with it's hierarchy
func (c *NodoCuentaContableCompositor) BuildTree() (treeData []*models.NodoArbolCuentaContable, err error) {
	rootNodes, _, err := nodoCcManager.GetRootNodes()

	if err != nil {
		return
	}
	_, noRootNodes, err := nodoCcManager.GetNoRootNodes()
	if err != nil {
		return
	}
	nodoCcHelper.BuildTreeFromDataSource(rootNodes, noRootNodes)
	treeData = rootNodes
	return
}
