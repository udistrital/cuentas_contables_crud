package migrations

import (
	"context"

	"github.com/udistrital/cuentas_contables_crud/models"

	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	CentroCostos := []interface{}{
		models.CentroCostos{
			ID:    "test_1",
			Label: "Test 1",
		},
		models.CentroCostos{
			ID:    "test_2",
			Label: "Test 2",
		},
	}
	migrate.Register(func(db *mongo.Database) error {

		_, err := db.Collection(models.CentroCostosCollection).InsertMany(context.TODO(), CentroCostos)
		if err != nil {
			return err
		}
		return nil
	}, func(db *mongo.Database) error {
		_, err := db.Collection(models.CentroCostosCollection).DeleteOne(context.TODO(), CentroCostos)
		if err != nil {
			return err
		}

		return nil
	})

}
