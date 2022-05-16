package migrations

import (
	"context"

	"github.com/udistrital/cuentas_contables_crud/models"

	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	naturalezaCuentaContable := []interface{}{
		models.NaturalezaCuentaContable{
			ID:    "credito",
			Label: "credito",
		},
		models.NaturalezaCuentaContable{
			ID:    "debito",
			Label: "debito",
		},
	}
	tipoMoneda := []interface{}{
		models.TipoMoneda{
			ID:    "cop",
			Label: "COP",
		},
	}
	migrate.Register(func(db *mongo.Database) error {

		err := db.Collection(models.NaturalezaCuentaContableCollection).Drop(context.TODO())
		if err != nil {
			return err
		}

		err = db.Collection(models.TipoMonedaCollection).Drop(context.TODO())
		if err != nil {
			return err
		}

		_, err = db.Collection(models.NaturalezaCuentaContableCollection).InsertMany(context.TODO(), naturalezaCuentaContable)
		if err != nil {
			return err
		}
		_, err = db.Collection(models.TipoMonedaCollection).InsertMany(context.TODO(), tipoMoneda)
		if err != nil {
			return err
		}
		return nil
	}, func(db *mongo.Database) error {

		_, err := db.Collection(models.NaturalezaCuentaContableCollection).DeleteOne(context.TODO(), naturalezaCuentaContable)
		if err != nil {
			return err
		}
		_, err = db.Collection(models.TipoMonedaCollection).DeleteOne(context.TODO(), tipoMoneda)
		if err != nil {
			return err
		}
		return nil
	})

}
