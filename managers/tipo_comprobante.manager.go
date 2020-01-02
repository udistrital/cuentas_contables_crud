package managers

import (
	"context"
	"errors"
	"time"

	"github.com/udistrital/cuentas_contables_crud/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TipoComprobanteManager this will manage the data process and store (CRUD) for the bussines (DAO)
type TipoComprobanteManager struct {
	Ctx context.Context
}

// AddItem This function will store the item data for the bussines proccess
func (m *TipoComprobanteManager) AddItem(itemData *models.TipoComprobante) (err error) {
	var tipoComprobante *models.TipoComprobante

	crudManager.Ctx = m.Ctx // Add ctx if process will be part of a transacction.

	if itemData.TipoDocumento != "" {
		_ = crudManager.GetDocumentByItem(itemData.TipoDocumento, "tipo_documento", models.TipoComprobanteCollection, &tipoComprobante)
	}
	if tipoComprobante != nil {
		return errors.New("item_exists")
	}
	general := models.General{
		FechaCreacion:     time.Now().Format("2006-01-02"),
		FechaModificacion: time.Now().Format("2006-01-02"),
		Activo:            true,
	}
	itemData.General = &general
	UUID, err := crudManager.AddDocument(itemData, models.TipoComprobanteCollection)

	if err != nil {
		return err
	}

	if UUID != "" {
		itemData.ID = UUID
	}

	return

}

// UpdateItem This function will store the item data for the bussines proccess
func (m *TipoComprobanteManager) UpdateItem(itemData *models.TipoComprobante, ID string) (err error) {
	var tipoComprobante *models.TipoComprobante
	var updtDoc interface{}
	crudManager.Ctx = m.Ctx // Add ctx if process will be part of a transacction.

	if itemData.TipoDocumento != "" {
		_ = crudManager.GetDocumentByItem(itemData.TipoDocumento, "tipo_documento", models.TipoComprobanteCollection, &tipoComprobante)
	}
	if tipoComprobante != nil {
		return errors.New("item_exists")
	}
	general := models.General{
		FechaCreacion:     itemData.FechaCreacion,
		FechaModificacion: time.Now().Format("2006-01-02"),
		Activo:            true,
	}
	itemData.General = &general
	objectID, _ := primitive.ObjectIDFromHex(ID)
	err = crudManager.UpdateDocument(itemData, objectID, models.TipoComprobanteCollection, updtDoc)

	if err != nil {
		return err
	}

	return

}
