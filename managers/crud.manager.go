package managers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/udistrital/cuentas_contables_crud/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CrudManager this manager must be used if you want to separate logic from implementation.
type CrudManager struct {
	Ctx context.Context
}

// GetDocumentByUUID get one document by it's uuid.
func (m *CrudManager) GetDocumentByUUID(UUID interface{}, collName string, resul interface{}) (err error) {
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

// GetDocumentByCodigo get one document by it's codigo.
func (m *CrudManager) GetDocumentByCodigo(codigo interface{}, collName string, resul interface{}) (err error) {
	coll, err := db.GetCollection(collName)

	if err != nil {
		return err
	}

	filter := make(map[string]interface{})

	filter["codigo"] = codigo

	err = coll.FindOne(context.TODO(), filter).Decode(resul)

	if err == mongo.ErrNoDocuments {
		return errors.New("document-no-found")
	}

	return
}

// GetDocumentByItem get one document by it's item by nameItem.
func (m *CrudManager) GetDocumentByItem(item interface{}, nameBson, collName string, resul interface{}) (err error) {
	coll, err := db.GetCollection(collName)

	if err != nil {
		return err
	}

	filter := make(map[string]interface{})

	filter[nameBson] = item

	err = coll.FindOne(context.TODO(), filter).Decode(resul)

	if err == mongo.ErrNoDocuments {
		return errors.New("document-no-found-by-item")
	}

	return
}

// DeleteDocumentByUUID delete one document by it's uuid.
func (m *CrudManager) DeleteDocumentByUUID(UUID interface{}, collName string, resul interface{}) (err error) {
	coll, err := db.GetCollection(collName)

	if err != nil {
		return err
	}

	filter := make(map[string]interface{})

	filter["_id"] = UUID

	resul, err = coll.DeleteOne(m.Ctx, filter)

	if err == mongo.ErrNoDocuments {
		return errors.New("cannot-delete-document")
	}

	return
}

// GetAllDocuments get one document by it's uuid.
func (m *CrudManager) GetAllDocuments(filter map[string]interface{}, limit, offset int64, collName string, fn func(*mongo.Cursor)) (err error) {
	coll, err := db.GetCollection(collName)
	if err != nil {
		return err
	}

	findOptions := options.Find()

	if limit >= 0 {
		findOptions.SetLimit(limit)
	}

	findOptions.SetSkip(offset)

	cur, err := coll.Find(context.TODO(), filter, findOptions)

	if err == mongo.ErrNoDocuments {
		return errors.New("document-no-found")
	} else if err != nil {
		log.Println("err", err.Error())
		return err
	}

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		fn(cur)

	}

	return
}

// UpdateDocument pdate one documen.
func (m *CrudManager) UpdateDocument(data interface{}, UUID interface{}, collName string, result interface{}) (err error) {

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

	go func() {

		update := map[string]interface{}{
			"$set": map[string]interface{}{
				"fecha_modificacion": time.Now().Format("2006-01-02"),
			},
		}

		res := coll.FindOneAndUpdate(m.Ctx, filter, update)

		if res.Err() != nil {
			log.Println("error:", res.Err().Error())
		}
	}()

	return
}

// AddDocument ...
func (m *CrudManager) AddDocument(data interface{}, collName string) (generatedID string, err error) {
	coll, err := db.GetCollection(collName)

	if err != nil {
		return "", err
	}

	resul, err := coll.InsertOne(m.Ctx, data)

	if err != nil {
		if strings.Contains(err.Error(), "dup key") {
			return "", errors.New("duplicated-document")
		}
		return "", err
	}
	generatedID, ok := resul.InsertedID.(string)
	if !ok {
		generatedID = resul.InsertedID.(primitive.ObjectID).Hex()
		if generatedID == "" {
			return "", errors.New("cannot-get-coll-id")
		}
	}

	go func() {
		filter := make(map[string]interface{})
		filter["_id"] = resul.InsertedID

		update := map[string]interface{}{
			"$set": map[string]interface{}{
				"fecha_creacion":     time.Now().Format("2006-01-02"),
				"fecha_modificacion": time.Now().Format("2006-01-02"),
				"activo":             true,
			},
		}
		res := coll.FindOneAndUpdate(m.Ctx, filter, update)

		if res.Err() != nil {
			log.Println("error:", res.Err().Error())
		}
	}()

	return
}

// RunTransaction ...
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
