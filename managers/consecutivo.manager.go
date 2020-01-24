package managers

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

// ConsecutivoManager this will manage the data process and store (CRUD) for the bussines (DAO)
type ConsecutivoManager struct {
	Ctx         context.Context
	crudManager CrudManager
}

// GetByCollection ...
func (m *ConsecutivoManager) GetByCollection(collectionName string) (data []map[string]interface{}) {
	filter := make(map[string]interface{})
	err := m.crudManager.GetAllDocuments(filter, -1, 0, collectionName, func(curr *mongo.Cursor) {
		var item map[string]interface{}
		if err := curr.Decode(&item); err == nil {
			data = append(data, item)
		}
	})
	if err != nil {
		errors.New("cannot-get-all-collection")
	}

	return
}
