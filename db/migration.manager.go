package db

import (
	"context"
	"time"

	"github.com/astaxie/beego"
	migrate "github.com/xakep666/mongo-migrate"

	// migrtions ... import the migration file.
	_ "github.com/udistrital/cuentas_contables_crud/migrations"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// RunMigrations ... Migrate all files in migrations package.
func RunMigrations() (*mongo.Database, error) {
	dbURL := beego.AppConfig.String("mongo_host")
	dbPort := beego.AppConfig.String("mongo_port")
	dbUser := beego.AppConfig.String("mongo_user")
	dbPass := beego.AppConfig.String("mongo_pass")
	dbAuth := beego.AppConfig.String("mongo_db_auth")
	dbMain := beego.AppConfig.String("mongo_db")
	clientOptions := options.Client().ApplyURI("mongodb://" + dbURL + ":" + dbPort).SetAuth(options.Credential{
		Username:   dbUser,
		Password:   dbPass,
		AuthSource: dbAuth, // db name
	})
	mainDB = dbMain
	var err error
	client, err = mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	db := client.Database(beego.AppConfig.String("mongo_db"))
	migrate.SetDatabase(db)
	if err := migrate.Up(migrate.AllAvailable); err != nil {
		return nil, err
	}
	return db, nil
}
