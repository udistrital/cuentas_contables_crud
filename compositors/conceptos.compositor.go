package compositors

import (
	"context"
	"github.com/udistrital/cuentas_contables_crud/managers"
	"github.com/udistrital/cuentas_contables_crud/helpers"
	"github.com/udistrital/cuentas_contables_crud/models"
)

// Conceptos compositor for controll the data processing over Conceptos models
type ConceptosCompositor struct {
	crudManager   managers.CrudManager
	nodoConcManager managers.ConceptosManager
	nodoConcHelper  helpers.ConceptosHelper
}

// AddNodeConceptos Add new node to the tree
func (c *ConceptosCompositor) AddNodeConceptos(nodeData *models.Conceptos) (err error) {

	err = c.crudManager.RunTransaction(func(ctx context.Context) error {
		concmang := managers.NewConceptosManager(nil)
		err = concmang.AddNodeConceptos(nodeData)
		return err
	})
	return
}

// BuildTree returns the tree data on the DB as a tree structure with it's hierarchy
func (c *ConceptosCompositor) BuildTreeConceptos(withNoActive ...bool) (rootNodes []*models.ArbolConceptosFormatNode, err error) {
	rootNodes, _, err = c.nodoConcManager.GetRootNodesConceptos(withNoActive...)
	if err != nil {
		return
	}
	_, noRootNodes, err := c.nodoConcManager.GetNoRootNodesConceptos(withNoActive...)
	if err != nil {
		return
	}
	c.nodoConcHelper.BuildTreeFromDataSource(rootNodes, noRootNodes)
	return
}

// GetNodeByCodigo Returns a *models.Conceptos by it's _id
func (c *ConceptosCompositor) GetNodeByCodigo(Codigo string) (node *models.Conceptos, err error) {

	resul := &models.Conceptos{}

	err = c.crudManager.GetDocumentByItem(Codigo, "codigo", models.ArbolConceptosCollection, resul)

	return resul, err
}
