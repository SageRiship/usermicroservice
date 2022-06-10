package connection

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var UserCollection *mongo.Collection
var Client *mongo.Client
var WalletCollection *mongo.Collection
var WalletTransactionCollection *mongo.Collection

var (
	Ctx = context.TODO()
	Db  *mongo.Database
)

func init() {
	host := "10.102.78.95"
	//host := "localhost"
	port := "27017"
	connectionURI := "mongodb://" + host + ":" + port + "/"
	clientOptions := options.Client().ApplyURI(connectionURI)
	Client, _ = mongo.Connect(Ctx, clientOptions)

	fmt.Println("Database : ", connectionURI)
	fmt.Println("MongoDB connected Successfully...")

	//Db = client.Database("UserData")              //Local database on PC
	Db = Client.Database("MVRDB")
	UserCollection = Db.Collection("User")
	WalletCollection = Db.Collection("Wallet")
	WalletTransactionCollection = Db.Collection("Wallet_transaction")

	//fmt.Println(UserCollection)

}
