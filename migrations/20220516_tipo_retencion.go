package migrations

import (
	"context"

	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/udistrital/cuentas_contables_crud/models"
)

func init() {
	tipoRetencion := []interface{}{
		models.TipoRetencion{
			ID:    "5",
			Label: "5",
		},
		models.TipoRetencion{
			ID:    "19",
			Label: "19",
		},
		models.TipoRetencion{
			ID:    "excento",
			Label: "Excento",
		},
		models.TipoRetencion{
			ID:    "excluido",
			Label: "Excluido",
		},
	}
	migrate.Register(func(db *mongo.Database) error {

		err := db.Collection(models.TipoRetencionCollection).Drop(context.TODO())
		if err != nil {
			return err
		}
		_, err = db.Collection(models.TipoRetencionCollection).InsertMany(context.TODO(), tipoRetencion)
		if err != nil {
			return err
		}
		return nil
	}, func(db *mongo.Database) error {
		_, err := db.Collection(models.TipoRetencionCollection).DeleteOne(context.TODO(), tipoRetencion)
		if err != nil {
			return err
		}

		return nil
	})

}
