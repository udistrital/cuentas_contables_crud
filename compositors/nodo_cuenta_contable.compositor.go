package compositors

import (
	"context"

	"github.com/udistrital/cuentas_contables_crud/managers"
	"github.com/udistrital/cuentas_contables_crud/models"
)

type NodoCuentaContableCompositor struct{}

var crudManager = managers.CrudManager{}

// GetNodeByID Returns a *models.NodoCuentaContable by it's _id
func (m *NodoCuentaContableCompositor) GetNodeByID(ID, collName string) (node *models.NodoCuentaContable, err error) {

	var resul *models.NodoCuentaContable

	err = crudManager.GetDocumentByUUID(ID, collName, resul)

	return resul, err
}

// AddNode Add new node to the tree
func (m *NodoCuentaContableCompositor) AddNode(nodeData *models.NodoCuentaContable, collName string) (err error) {

	err = crudManager.RunTransaction(func(ctx context.Context) error {
		ccmang := managers.NodoCuentaContableManager{
			// Ctx: ctx, // set this bar if mongo is deployed on replica set mode.
		}
		err = ccmang.AddNode(nodeData, collName)
		return err
	})
	return
}
