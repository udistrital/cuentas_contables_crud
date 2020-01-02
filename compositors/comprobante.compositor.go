package compositors

import (
	//"context"
	// "github.com/udistrital/cuentas_contables_crud/managers"
	"github.com/udistrital/cuentas_contables_crud/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ComprobanteCompositor struct{}

// GetComprobanteByID Returns a *models.Comprobante by it's _id
func (m *ComprobanteCompositor) GetComprobanteByID(ID string) (item *models.Comprobante, err error) {

	var resul *models.Comprobante
	objectID, _ := primitive.ObjectIDFromHex(ID)
	err = crudManager.GetDocumentByUUID(objectID, models.ComprobanteCollection, &resul)

	return resul, err
}

// GetAllComprobante Returns All Comprobante
func (m *ComprobanteCompositor) GetAllComprobante() (data []models.Comprobante, err error) {
	filter := make(map[string]interface{})
	dataIndexed := make(map[string]models.Comprobante)

	err = crudManager.GetAllDocuments(filter, -1, 0, models.ComprobanteCollection, func(curr *mongo.Cursor) {
		var item models.Comprobante
		if err := curr.Decode(&item); err == nil {
			dataIndexed[item.ID] = item
			data = append(data, item)
		}
	})

	return
}

// UpdateComprobante Update tipo_comprobante
func (m *ComprobanteCompositor) UpdateComprobante(itemData *models.Comprobante, ID string) (err error) {
	var updtDoc interface{}
	objectID, _ := primitive.ObjectIDFromHex(ID)
	err = crudManager.UpdateDocument(itemData, objectID, models.ComprobanteCollection, updtDoc)
	if err != nil {
		return err
	}
	return
}

// AddComprobante Add new node to the tree
func (m *ComprobanteCompositor) AddComprobante(nodeData *models.Comprobante) (err error) {

	_, err = crudManager.AddDocument(nodeData, models.ComprobanteCollection)
	if err != nil {
		return err
	}
	return
}

// DeleteComprobante Delete comprobante
func (m *ComprobanteCompositor) DeleteComprobante(ID string) (err error) {
	var updtDoc interface{}
	objectID, _ := primitive.ObjectIDFromHex(ID)
	err = crudManager.DeleteDocumentByUUID(objectID, models.ComprobanteCollection, updtDoc)
	if err != nil {
		return err
	}
	return
}
