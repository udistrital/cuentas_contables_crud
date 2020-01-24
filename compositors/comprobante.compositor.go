package compositors

import (
	//"context"
	// "github.com/udistrital/cuentas_contables_crud/managers"
	"context"

	"github.com/udistrital/cuentas_contables_crud/managers"
	"github.com/udistrital/cuentas_contables_crud/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ComprobanteCompositor ...
type ComprobanteCompositor struct {
	crudManager managers.CrudManager
}

// GetComprobanteByID Returns a *models.Comprobante by it's _id
func (m *ComprobanteCompositor) GetComprobanteByID(ID string) (item *models.Comprobante, err error) {

	var resul *models.Comprobante
	objectID, err := primitive.ObjectIDFromHex(ID)
	if errDocument := m.crudManager.GetDocumentByUUID(objectID, models.ComprobanteCollection, &resul); errDocument != nil {
		return resul, errDocument
	}

	return resul, err
}

// GetAllComprobante Returns All Comprobante
func (m *ComprobanteCompositor) GetAllComprobante() (data []models.Comprobante, err error) {
	filter := make(map[string]interface{})

	err = m.crudManager.GetAllDocuments(filter, -1, 0, models.ComprobanteCollection, func(curr *mongo.Cursor) {
		var item models.Comprobante
		if err := curr.Decode(&item); err == nil {
			data = append(data, item)
		}
	})

	return
}

// UpdateComprobante Update tipo_comprobante
func (m *ComprobanteCompositor) UpdateComprobante(itemData *models.Comprobante, ID string) (err error) {
	err = m.crudManager.RunTransaction(func(ctx context.Context) error {
		mang := managers.ComprobanteManager{
			// Ctx: ctx, // set this bar if mongo is deployed on replica set mode.
		}
		err = mang.UpdateItem(itemData, ID)
		return err
	})
	return
}

// AddComprobante Add new node to the tree
func (m *ComprobanteCompositor) AddComprobante(itemData *models.Comprobante) (err error) {

	err = m.crudManager.RunTransaction(func(ctx context.Context) error {
		mang := managers.ComprobanteManager{
			// Ctx: ctx, // set this bar if mongo is deployed on replica set mode.
		}
		err = mang.AddItem(itemData)
		return err
	})
	return
}

// DeleteComprobante Delete comprobante
func (m *ComprobanteCompositor) DeleteComprobante(ID string) (err error) {
	var updtDoc interface{}
	var objectID primitive.ObjectID
	if objectID, err = primitive.ObjectIDFromHex(ID); err != nil {
		return err
	}
	err = m.crudManager.DeleteDocumentByUUID(objectID, models.ComprobanteCollection, updtDoc)
	return
}
