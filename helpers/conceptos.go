package helpers

import (
	"github.com/udistrital/cuentas_contables_crud/models"
	"github.com/udistrital/utils_oas/errorctrl"
)

// ConceptosHelper ...
type ConceptosHelper struct{}

// BuildTreeFromDataSource ...
func (h *ConceptosHelper) BuildTreeFromDataSource(rootsData []*models.ArbolConceptosFormatNode, noRootNodesIndexed map[string]*models.NodoArbolConceptos) {
	defer errorctrl.ErrorControlFunction("BuildTreeFromDataSource", "500")
	for i := 0; i < len(rootsData); i++ {
		h.getTreeBranch(rootsData[i], noRootNodesIndexed)
	}
}

func (h *ConceptosHelper) getTreeBranch(nodeData *models.ArbolConceptosFormatNode, noRootNodesIndexed map[string]*models.NodoArbolConceptos) {
	defer errorctrl.ErrorControlFunction("getTreeBranch", "500")
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
