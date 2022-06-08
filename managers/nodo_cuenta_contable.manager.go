package managers

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/udistrital/cuentas_contables_crud/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// NodoCuentaContableManager this will manage the data process and store (CRUD) for the bussines (DAO)
// sthe collNme is necesary because we have two differents trees.
type NodoCuentaContableManager struct {
	Ctx         context.Context
	crudManager CrudManager
}

// NewNodoCuentaContableManager initialicer for this manager. useful if you want to pass the app context over the DB operations (transactions will need this)
func NewNodoCuentaContableManager(ctx context.Context) NodoCuentaContableManager {
	managerObj := NodoCuentaContableManager{
		Ctx: ctx,
		crudManager: CrudManager{
			Ctx: ctx,
		},
	}
	return managerObj
}

func (m *NodoCuentaContableManager) getNodesByFilter(filter map[string]interface{}, NaturalezaCuentaContable string, withNoActive ...bool) (nodesData []*models.NodoArbolCuentaContable, nodesDataIndexed map[string]*models.NodoArbolCuentaContable, err error) {
	localfilter := make(map[string]interface{})
	if filter != nil {
		localfilter = filter
	}

	if len(withNoActive) == 0 || (len(withNoActive) > 0 && !withNoActive[0]) {
		localfilter["activo"] = true
	}

	nodesDataIndexed = make(map[string]*models.NodoArbolCuentaContable)

	if NaturalezaCuentaContable != "" {
		localfilter["naturaleza_id"] = NaturalezaCuentaContable
	}
	err = m.crudManager.GetAllDocuments(filter, -1, 0, models.ArbolPlanMaestroCuentasContCollection, func(curr *mongo.Cursor) {
		var node models.NodoArbolCuentaContable
		if err := curr.Decode(&node); err == nil {
			nodesDataIndexed[node.Codigo] = &node
			nodesData = append(nodesData, &node)
		}
	})

	return
}

// AddNode This function will store the node data of a tree for the bussines proccess
func (m *NodoCuentaContableManager) AddNode(nodeData *models.NodoCuentaContable) (err error) {
	var fatherData *models.NodoCuentaContable
	var tempResults interface{}

	if e := m.crudManager.GetDocumentByUUID(nodeData.NaturalezaCuentaID, models.NaturalezaCuentaContableCollection, &tempResults); e != nil {
		return errors.New("naturaleza-no-found")
	}

	if e := m.crudManager.GetDocumentByUUID(nodeData.MonedaID, models.TipoMonedaCollection, &tempResults); e != nil {
		return errors.New("tipo-moneda-no-found")
	}

	if nodeData.Padre != nil {
		if e := m.crudManager.GetDocumentByCodigo(*nodeData.Padre, models.ArbolPlanMaestroCuentasContCollection, &fatherData); e != nil {
			return errors.New("father-no-found")
		}
	}

	nodeData.General = &models.General{}
	nodeData.Activo = true
	originalID := nodeData.Codigo
	if fatherData != nil { // infer level from father if it exist.
		nodeData.Nivel = fatherData.Nivel + 1
		nodeData.Codigo = fatherData.Codigo + "-" + nodeData.Codigo
	} else {
		nodeData.Nivel = 1 // put 1 as default level
	}
	// check for curr level constraints.
	if currLevelParam, e := m.GetLevelParameterForNode(nodeData.Nivel); e == nil {
		if len(originalID) != *currLevelParam.CodeLenght {
			return errors.New("code-lenght-error")
		}
	} else {
		log.Println("error", e.Error())
		return errors.New("parameter-for-level-no-found")
	}
	UUID, err := m.crudManager.AddDocument(nodeData, models.ArbolPlanMaestroCuentasContCollection)

	if err != nil {
		return err
	}

	if UUID != "" {
		nodeData.ID, _ = primitive.ObjectIDFromHex(UUID)
	}

	if fatherData != nil {
		fatherData.Hijos = append(fatherData.Hijos, nodeData.Codigo)
		var updtDoc interface{}
		updMap := map[string]interface{}{
			"hijos": fatherData.Hijos,
		}
		if e := m.crudManager.UpdateDocument(updMap, fatherData.ID, models.ArbolPlanMaestroCuentasContCollection, updtDoc); e != nil {
			return e
		}
	}

	return

}

// GetRootNodes returns the "Plan maestro" tree's root nodes
func (m *NodoCuentaContableManager) GetRootNodes(NaturalezaCuentaContable string, withNoActive ...bool) (rootsDataFormated []*models.ArbolNbFormatNode, nodesDataIndexed map[string]*models.NodoArbolCuentaContable, err error) {
	var codigo []string
	filter := make(map[string]interface{})

	if NaturalezaCuentaContable != "" {
		filter = map[string]interface{}{"naturaleza_id": NaturalezaCuentaContable}
	} else {
		filter = map[string]interface{}{"padre": nil}
	}

	rootsData, nodesDataIndexed, err := m.getNodesByFilter(filter, NaturalezaCuentaContable, withNoActive...)

	if err != nil {
		return
	}

	for _, root := range rootsData {
		index := strings.IndexAny(root.Codigo, "-")
		if index == -1 {
			index = len(root.Codigo)
		}
		newdata := stringInSlice(root.Codigo[0:index], codigo)
		if newdata {
			continue
		}
		codigo = append(codigo, root.Codigo[0:1])
		rootsDataFormated = append(rootsDataFormated, &models.ArbolNbFormatNode{
			Data: root,
		})
	}
	return
}

// GetNoRootNodes returns the "Plan maestro" tree's non root nodes
func (m *NodoCuentaContableManager) GetNoRootNodes(NaturalezaCuentaContable string, withNoActive ...bool) (nodesData []*models.NodoArbolCuentaContable, nodesDataIndexed map[string]*models.NodoArbolCuentaContable, err error) {

	filter := map[string]interface{}{"padre": map[string]interface{}{"$ne": nil}}

	nodesData, nodesDataIndexed, err = m.getNodesByFilter(filter, NaturalezaCuentaContable, withNoActive...)
	return
}

// ChangeNodeState this function will enable or disable one node from the tree (if a root node is disabled, full branch will no be visible in some services)
func (m *NodoCuentaContableManager) ChangeNodeState(UUID string) (err error) {

	var nodeData models.NodoCuentaContable
	var result interface{}

	err = m.crudManager.GetDocumentByUUID(UUID, models.ArbolPlanMaestroCuentasContCollection, &nodeData)

	if err != nil {
		return
	}

	updateData := map[string]interface{}{
		"activo": !nodeData.Activo,
	}
	err = m.crudManager.UpdateDocument(updateData, UUID, models.ArbolPlanMaestroCuentasContCollection, &result)

	return
}

// GetLevelParameterForNode returns the parameter value of a specific level for the plan cuentas tree or error "parameter-no-found".
func (m *NodoCuentaContableManager) GetLevelParameterForNode(level int) (*models.ArbolCuentaContableParameters, error) {
	filter := map[string]interface{}{
		"nivel": level,
	}
	var parameter *models.ArbolCuentaContableParameters
	err := m.crudManager.GetAllDocuments(filter, 1, 0, models.ArbolCuentasContParametersCollection, func(curr *mongo.Cursor) {

		if err := curr.Decode(&parameter); err == nil {
		}
		return
	})

	return parameter, err
}

// stringInSlice returns true/false if there is a repeated root node
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// deleteNodeByUUID this function will delete a node from the tree
func (m *NodoCuentaContableManager) DeleteNodeByUUID(id interface{}) (err error) {

	var nodeData models.NodoCuentaContable
	var result interface{}

	err = m.crudManager.GetDocumentByUUID(id, models.ArbolPlanMaestroCuentasContCollection, &nodeData)

	if err != nil {
		return
	}

	if len(nodeData.Hijos) > 0 {
		return errors.New("node-has-children")
	}

	err = m.crudManager.DeleteDocumentByUUID(id, models.ArbolPlanMaestroCuentasContCollection, &result)

	return
}
