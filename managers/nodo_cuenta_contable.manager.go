package managers

import (
	"context"

	"github.com/udistrital/cuentas_contables_crud/models"
)

// NodoCuentaContableManager this will manage the data process and store (CRUD) for the bussines (DAO)
// sthe collNme is necesary because we have two differents trees.
type NodoCuentaContableManager struct {
	Ctx context.Context
}

var crudManager = CrudManager{}

// AddNode This function will store the node data of a tree for the bussines proccess
func (m *NodoCuentaContableManager) AddNode(nodeData *models.NodoCuentaContable, collName string) (err error) {
	var fatherData *models.NodoCuentaContable

	crudManager.Ctx = m.Ctx // Add ctx if process will be part of a transacction.

	if nodeData.Padre != "" {
		if e := crudManager.GetDocumentByUUID(nodeData.Padre, collName, &fatherData); e != nil {
			return e
		}
	}

	UUID, err := crudManager.AddDocument(nodeData, collName)

	if err != nil {
		return err
	}

	if UUID != "" {
		nodeData.ID = UUID
	}

	if fatherData != nil {
		fatherData.Hijos = append(fatherData.Hijos, nodeData.ID)
		var updtDoc interface{}
		updMap := map[string]interface{}{
			"hijos": fatherData.Hijos,
		}
		if e := crudManager.UpdateDocument(updMap, fatherData.ID, collName, updtDoc); e != nil {
			return e
		}
	}

	return

}
