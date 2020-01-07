package migrations

import (
	"context"

	"github.com/udistrital/cuentas_contables_crud/models"

	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {

	level1Obj := models.ArbolCuentaContableParameters{}
	level1Level := 1
	level1Lenght := 1
	level1Obj.Nivel = &level1Level
	level1Obj.CodeLenght = &level1Lenght
	//-----------------
	level2Obj := models.ArbolCuentaContableParameters{}
	level2Level := 2
	level2Lenght := 1
	level2Obj.Nivel = &level2Level
	level2Obj.CodeLenght = &level2Lenght
	//-----------------
	level3Obj := models.ArbolCuentaContableParameters{}
	level3Level := 3
	level3Lenght := 2
	level3Obj.Nivel = &level3Level
	level3Obj.CodeLenght = &level3Lenght
	//-----------------
	level4Obj := models.ArbolCuentaContableParameters{}
	level4Level := 4
	level4Lenght := 2
	level4Obj.Nivel = &level4Level
	level4Obj.CodeLenght = &level4Lenght
	//-----------------
	level5Obj := models.ArbolCuentaContableParameters{}
	level5Level := 5
	level5Lenght := 2
	level5Obj.Nivel = &level5Level
	level5Obj.CodeLenght = &level5Lenght
	//-----------------
	level6Obj := models.ArbolCuentaContableParameters{}
	level6Level := 5
	level6Lenght := 2
	level6Obj.Nivel = &level6Level
	level6Obj.CodeLenght = &level6Lenght
	parameters := []interface{}{
		level1Obj,
		level2Obj,
		level3Obj,
		level4Obj,
		level5Obj,
		level6Obj,
	}

	migrate.Register(func(db *mongo.Database) error {

		_, err := db.Collection(models.ArbolCuentasContParametersCollection).InsertMany(context.TODO(), parameters)
		if err != nil {
			return err
		}

		return nil
	}, func(db *mongo.Database) error {
		_, err := db.Collection(models.ArbolCuentasContParametersCollection).DeleteOne(context.TODO(), parameters)
		if err != nil {
			return err
		}

		return nil
	})

}
