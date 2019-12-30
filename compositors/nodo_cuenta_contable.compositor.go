package compositors

import (
	"context"

	"github.com/udistrital/cuentas_contables_crud/managers"
	"github.com/udistrital/cuentas_contables_crud/models"
)

type NodoCuentaContableCompositor struct{}

var crudManager = managers.CrudManager{}

// GetNodeByID Returns a *models.NodoCuentaContable by it's _id
func (m *NodoCuentaContableCompositor) GetNodeByID(ID string) (node *models.NodoCuentaContable, err error) {

	resul := &models.NodoCuentaContable{}

	err = crudManager.GetDocumentByUUID(ID, models.ArbolPlanMaestroCuentasContCollection, resul)

	return resul, err
}

// AddNode Add new node to the tree
func (c *NodoCuentaContableCompositor) AddNode(nodeData *models.NodoCuentaContable) (err error) {

	err = crudManager.RunTransaction(func(ctx context.Context) error {
		ccmang := managers.NodoCuentaContableManager{
			// Ctx: ctx, // set this var if mongo is deployed on replica set mode.
		}
		err = ccmang.AddNode(nodeData)
		return err
	})
	return
}

func (c *NodoCuentaContableCompositor) BuildTree() (err error) {
	return nil
}
