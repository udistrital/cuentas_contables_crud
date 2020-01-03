package helpers

import "github.com/udistrital/cuentas_contables_crud/models"

type NodoCuentaContableHelper struct{}

func (h *NodoCuentaContableHelper) BuildTreeFromDataSource(rootsData []*models.NodoArbolCuentaContable, noRootNodesIndexed map[string]*models.NodoArbolCuentaContable) {
	for i := 0; i < len(rootsData); i++ {
		h.getTreeBranch(rootsData[i], noRootNodesIndexed)
	}
}

func (h *NodoCuentaContableHelper) getTreeBranch(nodeData *models.NodoArbolCuentaContable, noRootNodesIndexed map[string]*models.NodoArbolCuentaContable) {
	for _, childRef := range nodeData.Hijos {
		if noRootNodesIndexed[childRef] != nil {
			h.getTreeBranch(noRootNodesIndexed[childRef], noRootNodesIndexed)
			nodeData.HijosRef = append(nodeData.HijosRef, noRootNodesIndexed[childRef])
			delete(noRootNodesIndexed, childRef)
		}
	}

}
