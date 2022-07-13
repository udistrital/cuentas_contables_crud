package migrations

import (
	"context"

	"github.com/udistrital/cuentas_contables_crud/models"

	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	tipoCuentaCOntable := []interface{}{
		models.DetalleCuentaContable{
			ID:    "no_detallada",
			Label: "No Detallada",
		},
		models.DetalleCuentaContable{
			ID:    "por_pagar",
			Label: "Por Pagar",
		},
		models.DetalleCuentaContable{
			ID:    "por_cobrar",
			Label: "Por Cobrar",
		},
		models.DetalleCuentaContable{
			ID:    "iva",
			Label: "IVA",
		},
		models.DetalleCuentaContable{
			ID:    "reteiva",
			Label: "ReteIVA",
		},
		models.DetalleCuentaContable{
			ID:    "bancos",
			Label: "Bancos",
		},
		models.DetalleCuentaContable{
			ID:    "gravada",
			Label: "Gravada",
		},

		models.DetalleCuentaContable{
			ID:    "no_gravada",
			Label: "No gravada",
		},

		models.DetalleCuentaContable{
			ID:    "exportacion",
			Label: "Exportación",
		},
		models.DetalleCuentaContable{
			ID:    "importacion_gravada",
			Label: "Importación gravada",
		},
		models.DetalleCuentaContable{
			ID:    "importacion_no_gravada",
			Label: "Importación no gravada",
		},
		models.DetalleCuentaContable{
			ID:    "inventario",
			Label: "Inventario",
		},
	}
	migrate.Register(func(db *mongo.Database) error {

		err := db.Collection(models.DetalleCuentaContableCollection).Drop(context.TODO())
		if err != nil {
			return err
		}

		_, err = db.Collection(models.DetalleCuentaContableCollection).InsertMany(context.TODO(), tipoCuentaCOntable)
		if err != nil {
			return err
		}
		return nil
	}, func(db *mongo.Database) error {
		_, err := db.Collection(models.DetalleCuentaContableCollection).DeleteOne(context.TODO(), tipoCuentaCOntable)
		if err != nil {
			return err
		}

		return nil
	})

}
