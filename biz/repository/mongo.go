package repository

import (
	"context"
	"fmt"
	"github.com/lutasam/doctors/biz/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoDB *mongo.Client

func init() {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s",
		utils.GetConfigString("mongo.address")),
	)
	var err error
	mongoDB, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}
}

func GetMongo() *mongo.Client {
	return mongoDB
}
