package compositors

import (
	"context"

	"github.com/udistrital/cuentas_contables_crud/managers"
	"github.com/udistrital/cuentas_contables_crud/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TipoComprobanteCompositor struct {
	crudManager managers.CrudManager
}

// var crudManager = managers.CrudManager{}

// GetAllTipoComprobante Returns All TipoComprobante
func (m *TipoComprobanteCompositor) GetAllTipoComprobante() (data []models.TipoComprobante, err error) {
	filter := make(map[string]interface{})
	dataIndexed := make(map[string]models.TipoComprobante)

	err = m.crudManager.GetAllDocuments(filter, -1, 0, models.TipoComprobanteCollection, func(curr *mongo.Cursor) {
		var item models.TipoComprobante
		if err := curr.Decode(&item); err == nil {
			dataIndexed[item.ID] = item
			data = append(data, item)
		}
	})

	return
}

// GetTipoComprobanteByID Returns a *models.TipoComprobante by it's _id
func (m *TipoComprobanteCompositor) GetTipoComprobanteByID(ID string) (item *models.TipoComprobante, err error) {

	var resul *models.TipoComprobante
	objectID, _ := primitive.ObjectIDFromHex(ID)

	err = m.crudManager.GetDocumentByItem(objectID, "_id", models.TipoComprobanteCollection, &resul)

	return resul, err
}

// AddTipoComprobante Add new tipo_comprobante
func (m *TipoComprobanteCompositor) AddTipoComprobante(itemData *models.TipoComprobante) (err error) {

	err = m.crudManager.RunTransaction(func(ctx context.Context) error {
		mang := managers.TipoComprobanteManager{
			// Ctx: ctx, // set this bar if mongo is deployed on replica set mode.
		}
		err = mang.AddItem(itemData)
		return err
	})
	return
}

// UpdateTipoComprobante Update tipo_comprobante
func (m *TipoComprobanteCompositor) UpdateTipoComprobante(itemData *models.TipoComprobante, ID string) (err error) {
	err = m.crudManager.RunTransaction(func(ctx context.Context) error {
		mang := managers.TipoComprobanteManager{
			// Ctx: ctx, // set this bar if mongo is deployed on replica set mode.
		}
		err = mang.UpdateItem(itemData, ID)
		return err
	})
	return
}

// DeleteTipoComprobante Delete tipo_comprobante
func (m *TipoComprobanteCompositor) DeleteTipoComprobante(ID string) (err error) {
	var updtDoc interface{}
	objectID, _ := primitive.ObjectIDFromHex(ID)
	err = m.crudManager.DeleteDocumentByUUID(objectID, models.TipoComprobanteCollection, updtDoc)
	if err != nil {
		return err
	}
	return
}
