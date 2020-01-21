package managers

import (
	"context"
	"errors"
	"time"

	"github.com/udistrital/cuentas_contables_crud/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ComprobanteManager this will manage the data process and store (CRUD) for the bussines (DAO)
type ComprobanteManager struct {
	Ctx                context.Context
	crudManager        CrudManager
	consecutivoManager ConsecutivoManager
}

// AddItem This function will store the item data for the bussines proccess
func (m *ComprobanteManager) AddItem(itemData *models.Comprobante) (err error) {
	m.crudManager.Ctx = m.Ctx                                                                             // Add ctx if process will be part of a transacction.
	consecutivoComprobante := len(m.consecutivoManager.GetByCollection(models.ComprobanteCollection)) + 1 //TODO: process consecutive CORE
	general := models.General{
		FechaCreacion:     time.Now().Format("2006-01-02"),
		FechaModificacion: time.Now().Format("2006-01-02"),
		Activo:            true,
	}
	itemData.Codigo = consecutivoComprobante
	itemData.General = &general
	if itemData.TipoComprobante == nil {
		return errors.New("TipoComprobante Not Defined")
	}
	UUID, err := m.crudManager.AddDocument(itemData, models.ComprobanteCollection)

	if err != nil {
		return err
	}

	if UUID != "" {
		itemData.ID = UUID
	}

	return

}

// UpdateItem This function will store the item data for the bussines proccess
func (m *ComprobanteManager) UpdateItem(itemData *models.Comprobante, ID string) (err error) {
	var updtDoc interface{}
	m.crudManager.Ctx = m.Ctx // Add ctx if process will be part of a transacction.

	general := models.General{
		FechaCreacion:     itemData.FechaCreacion,
		FechaModificacion: time.Now().Format("2006-01-02"),
		Activo:            true,
	}
	itemData.General = &general
	itemData.ID = ""
	objectID, _ := primitive.ObjectIDFromHex(ID)
	err = m.crudManager.UpdateDocument(itemData, objectID, models.ComprobanteCollection, updtDoc)

	if err != nil {
		return err
	}

	return

}
