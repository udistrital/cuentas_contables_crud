package helpers

import "github.com/udistrital/cuentas_contables_crud/models"

// NodoCuentaContableHelper ...
type NodoCuentaContableHelper struct{}

// BuildTreeFromDataSource ...
func (h *NodoCuentaContableHelper) BuildTreeFromDataSource(rootsData []*models.ArbolNbFormatNode, noRootNodesIndexed map[string]*models.NodoArbolCuentaContable) {
	for i := 0; i < len(rootsData); i++ {
		h.getTreeBranch(rootsData[i], noRootNodesIndexed)
	}
}

func (h *NodoCuentaContableHelper) getTreeBranch(nodeData *models.ArbolNbFormatNode, noRootNodesIndexed map[string]*models.NodoArbolCuentaContable) {
	for _, childRef := range nodeData.Data.Hijos {
		if noRootNodesIndexed[childRef] != nil {
			nodoFormated := &models.ArbolNbFormatNode{
				Data: noRootNodesIndexed[childRef],
			}
			h.getTreeBranch(nodoFormated, noRootNodesIndexed)
			nodeData.Children = append(nodeData.Children, nodoFormated)
			delete(noRootNodesIndexed, childRef)
		}
	}

}
