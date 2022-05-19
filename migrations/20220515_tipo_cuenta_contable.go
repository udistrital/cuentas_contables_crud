package migrations

import (
	"context"

	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/udistrital/cuentas_contables_crud/models"
)

func init() {
	tipoCuentaContable := []interface{}{
		models.TipoCuenta{
			ID:    "activo",
			Label: "Activo",
		},
		models.TipoCuenta{
			ID:    "pasivo",
			Label: "Pasivo",
		},
		models.TipoCuenta{
			ID:    "patrimonio",
			Label: "Patrimonio",
		},
		models.TipoCuenta{
			ID:    "ingreso",
			Label: "Ingresos",
		},
		models.TipoCuenta{
			ID:    "gasto",
			Label: "Gastos",
		},
		models.TipoCuenta{
			ID:    "costo_venta",
			Label: "Constos de Ventas",
		},
		models.TipoCuenta{
			ID:    "costo_produccion",
			Label: "Constos de produccion o de operacion",
		},
		models.TipoCuenta{
			ID:    "costo_orden_deudor",
			Label: "Cuentas de orden deudoras",
		},
		models.TipoCuenta{
			ID:    "costo_orden_acreedor",
			Label: "Cuentas de orden acreedoras",
		},
	}
	migrate.Register(func(db *mongo.Database) error {

		err := db.Collection(models.TipoCuentaCollection).Drop(context.TODO())
		if err != nil {
			return err
		}

		_, err = db.Collection(models.TipoCuentaCollection).InsertMany(context.TODO(), tipoCuentaContable)
		if err != nil {
			return err
		}
		return nil
	}, func(db *mongo.Database) error {
		_, err := db.Collection(models.TipoCuentaCollection).DeleteOne(context.TODO(), tipoCuentaContable)
		if err != nil {
			return err
		}
		return nil
	})

}
