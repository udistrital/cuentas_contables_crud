package managers

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/udistrital/cuentas_contables_crud/db"
	"go.mongodb.org/mongo-driver/mongo"
)

// CrudManager this manager must be used if you want to separate logic from implementation.
type CrudManager struct {
	Ctx context.Context
}

// GetDocumentByUUID get one document by it's uuid.
func (m *CrudManager) GetDocumentByUUID(UUID, collName string, resul interface{}) (err error) {
	coll, err := db.GetCollection(collName)

	if err != nil {
		return err
	}

	filter := make(map[string]interface{})

	filter["_id"] = UUID

	err = coll.FindOne(context.TODO(), filter).Decode(resul)

	if err == mongo.ErrNoDocuments {
		return errors.New("document-no-found")
	}

	return
}

// UpdateDocument pdate one documen.
func (m *CrudManager) UpdateDocument(data interface{}, UUID, collName string, result interface{}) (err error) {

	coll, err := db.GetCollection(collName)

	if err != nil {
		return err
	}
	filter := make(map[string]interface{})

	filter["_id"] = UUID

	update := map[string]interface{}{
		"$set": data,
	}

	res := coll.FindOneAndUpdate(m.Ctx, filter, update)

	if res.Err() != nil {
		log.Println("error:", res.Err().Error())
		return errors.New("cannot-update-document")
	}

	return
}

func (m *CrudManager) AddDocument(data interface{}, collName string) (generatedID string, err error) {
	coll, err := db.GetCollection(collName)

	if err != nil {
		return "", err
	}

	resul, err := coll.InsertOne(m.Ctx, data)

	if err != nil {
		return "", err
	}

	generatedID, ok := resul.InsertedID.(string)

	if !ok {
		return "", errors.New("cannot-get-coll-id")
	}

	return
}

func (m *CrudManager) RunTransaction(f func(context.Context) error) (err error) {

	client, err := db.GetClient()

	if err != nil {
		return
	}

	session, err := client.StartSession()

	if err != nil {
		return
	}

	if err = session.StartTransaction(); err != nil {
		return
	}

	if e := mongo.WithSession(context.Background(), session, func(sc mongo.SessionContext) error {
		defer func() {
			if r := recover(); r != nil {

				err = session.AbortTransaction(sc)
				if err == nil {
					fmt.Println("error: ", r)
					err = fmt.Errorf("internal-server-error")
				}
			}
		}()

		err = f(sc)

		if err != nil {
			eTr := session.AbortTransaction(sc)
			if eTr == nil {
				return err
			}
			return eTr

		}

		err = session.CommitTransaction(sc)
		return err
	}); e != nil {
		return e
	}
	session.EndSession(context.Background())

	return
}
