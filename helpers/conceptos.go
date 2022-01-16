package helpers

import (
	"github.com/udistrital/cuentas_contables_crud/models"
)

// ConceptosHelper ...
type ConceptosHelper struct{}

// BuildTreeFromDataSource ...
func (h *ConceptosHelper) BuildTreeFromDataSource(rootsData []*models.ArbolConceptosFormatNode, noRootNodesIndexed map[string]*models.NodoArbolConceptos) {
	for i := 0; i < len(rootsData); i++ {
		h.getTreeBranch(rootsData[i], noRootNodesIndexed)
	}
}

func (h *ConceptosHelper) getTreeBranch(nodeData *models.ArbolConceptosFormatNode, noRootNodesIndexed map[string]*models.NodoArbolConceptos) {
	for _, childRef := range nodeData.Data.Hijos {
		if noRootNodesIndexed[childRef] != nil {
			nodoFormated := &models.ArbolConceptosFormatNode{
				Data: noRootNodesIndexed[childRef],
			}
			h.getTreeBranch(nodoFormated, noRootNodesIndexed)
			nodeData.Children = append(nodeData.Children, nodoFormated)
			delete(noRootNodesIndexed, childRef)
		}
	}

}
