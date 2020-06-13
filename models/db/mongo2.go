package db

import (
	"context"
	"log"

	"github.com/hepiska/todo-go/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//ConfigDB get mongo conectrion
func ConfigDB() *mongo.Database {
	connectionString := utils.EnvVar("DB_CONNECTION_STRING")
	dbName := utils.EnvVar("DB_NAME")
	ctx := context.Background()
	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(connectionString),
	)
	if err != nil {
		log.Fatal(err)
	}
	return client.Database(dbName)

}
