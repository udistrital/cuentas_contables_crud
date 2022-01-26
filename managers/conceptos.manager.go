package managers

import (
	"context"
	"errors"
	_ "log"
	"strings"

	"github.com/udistrital/cuentas_contables_crud/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ConceptosManager struct {
	Ctx         context.Context
	crudManager CrudManager
}

// NewConceptosManager initialicer for this manager. useful if you want to pass the app context over the DB operations (transactions will need this)
func NewConceptosManager(ctx context.Context) ConceptosManager {
	managerObj := ConceptosManager{
		Ctx: ctx,
		crudManager: CrudManager{
			Ctx: ctx,
		},
	}
	return managerObj
}

// AddNodeConceptos This function will store the node data of a tree for the bussines proccess
func (m *ConceptosManager) AddNodeConceptos(nodeData *models.Conceptos) (err error) {
	var fatherData *models.Conceptos
	var tempResults interface{}
	var objectID primitive.ObjectID
	if objectID, err = primitive.ObjectIDFromHex(nodeData.TipoComprobanteId); err != nil {
		return err
	}
	if e := m.crudManager.GetDocumentByUUID(objectID, models.TipoComprobanteCollection, &tempResults); e != nil {
		return errors.New("tipo-comprobante-no-found")
	}

	if nodeData.Padre != nil {
		if e := m.crudManager.GetDocumentByItem(*nodeData.Padre, "codigo", models.ArbolConceptosCollection, &fatherData); e != nil {
			return errors.New("father-no-found")
		}
	}
	nodeData.ID = primitive.NewObjectID()
	nodeData.General = &models.General{}
	nodeData.Activo = true
	//originalID := nodeData.Codigo
	if fatherData != nil { // infer level from father if it exist.
		nodeData.Nivel = fatherData.Nivel + 1
		nodeData.Codigo = fatherData.Codigo + "-" + nodeData.Codigo
	} else {
		nodeData.Nivel = 1 // put 1 as default level
	}
	var tempResults2 interface{}
	if _ = m.crudManager.GetDocumentByItem(nodeData.Codigo, "codigo", models.ArbolConceptosCollection, &tempResults2); tempResults2 != nil {
		return errors.New("code-already-exists")
	}
	_, err = m.crudManager.AddDocument(nodeData, models.ArbolConceptosCollection)
	if err != nil {
		return err
	}

	if fatherData != nil {
		fatherData.Hijos = append(fatherData.Hijos, nodeData.ID.Hex())
		var updtDoc interface{}
		updMap := map[string]interface{}{
			"hijos": fatherData.Hijos,
		}
		if e := m.crudManager.UpdateDocument(updMap, fatherData.ID, models.ArbolConceptosCollection, updtDoc); e != nil {
			return e
		}
	}
	return
}

// GetLevelParameterForNode returns the parameter value of a specific level for the plan cuentas tree or error "parameter-no-found".
func (m *ConceptosManager) GetLevelParameterForNodeConceptos(level int) (*models.ArbolConceptosParameters, error) {
	filter := map[string]interface{}{
		"nivel": level,
	}
	var parameter *models.ArbolConceptosParameters
	err := m.crudManager.GetAllDocuments(filter, 1, 0, models.ArbolConceptosParametersCollection, func(curr *mongo.Cursor) {
		curr.Decode(&parameter)
		return
	})

	return parameter, err
}

// GetRootNodesConceptos returns the "Plan maestro" tree's root nodes
func (m *ConceptosManager) GetRootNodesConceptos(withNoActive ...bool) (rootsDataFormated []*models.ArbolConceptosFormatNode, nodesDataIndexed map[string]*models.NodoArbolConceptos, err error) {
	var codigo []string
	var rootsData []*models.NodoArbolConceptos
	filter := make(map[string]interface{})

	filter = map[string]interface{}{"padre": nil}

	rootsData, nodesDataIndexed, err = m.getNodesByFilterConceptos(filter, withNoActive...)
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
		rootsDataFormated = append(rootsDataFormated, &models.ArbolConceptosFormatNode{
			Data: root,
		})
	}
	return
}

// GetNoRootNodesConceptos returns the "Plan maestro" tree's non root nodes
func (m *ConceptosManager) GetNoRootNodesConceptos(withNoActive ...bool) (nodesData []*models.NodoArbolConceptos, nodesDataIndexed map[string]*models.NodoArbolConceptos, err error) {

	filter := map[string]interface{}{"padre": nil}
	nodesData, nodesDataIndexed, err = m.getNodesByFilterConceptos(filter, withNoActive...)
	return
}

func (m *ConceptosManager) getNodesByFilterConceptos(filter map[string]interface{}, withNoActive ...bool) (nodesData []*models.NodoArbolConceptos, nodesDataIndexed map[string]*models.NodoArbolConceptos, err error) {

	if len(withNoActive) == 0 || (len(withNoActive) > 0 && !withNoActive[0]) {
		filter["activo"] = true
	}

	nodesDataIndexed = make(map[string]*models.NodoArbolConceptos)

	err = m.crudManager.GetAllDocuments(filter, -1, 0, models.ArbolConceptosCollection, func(curr *mongo.Cursor) {
		var node models.NodoArbolConceptos
		if err := curr.Decode(&node); err == nil {
			nodesDataIndexed[node.ID.Hex()] = &node
			nodesData = append(nodesData, &node)
		}
	})

	return
}
