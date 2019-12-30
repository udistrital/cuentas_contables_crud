package compositors

import (
	"context"

	"github.com/udistrital/cuentas_contables_crud/managers"
	"github.com/udistrital/cuentas_contables_crud/models"
	"go.mongodb.org/mongo-driver/mongo"
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

func (c *NodoCuentaContableCompositor) BuildTree() (treeData []models.NodoArbolCuentaContable, err error) {
	filter := make(map[string]interface{})
	dataIndexed := make(map[string]models.NodoArbolCuentaContable)

	err = crudManager.GetAllDocuments(filter, -1, 0, models.ArbolPlanMaestroCuentasContCollection, func(curr *mongo.Cursor) {
		var node models.NodoArbolCuentaContable
		if err := curr.Decode(&node); err == nil {
			dataIndexed[node.ID] = node
			treeData = append(treeData, node)
		}
	})

	return
}
