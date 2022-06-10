package service

import (
	"context"
	"fmt"
	"log"
	"time"

	coll "github.com/SageRiship/usermicroservice/connection"
	"github.com/SageRiship/usermicroservice/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddWalletService(wallet model.Wallet) (*model.Wallet, error) {
	/*
		if we want to generate id by GO write this and
			[ Id          primitive.ObjectID  `json:"_id" bson:"_id"` ]..in struct
	*/
	wallet.Id = primitive.NewObjectID()
	wallet.CreatedOn = primitive.NewDateTimeFromTime(time.Now())
	wallet.UpdatedOn = primitive.NewDateTimeFromTime(time.Now())

	inserted, err := coll.WalletCollection.InsertOne(context.Background(), wallet)

	if err != nil {

		return nil, err
	}
	fmt.Println("Inserted 1 Wallet in db with id: ", inserted.InsertedID)
	return &wallet, nil

}

func GetAllWalletService() ([]model.Wallet, error) {

	var walletList []model.Wallet
	cursor, err := coll.WalletCollection.Find(context.Background(), bson.D{})
	if cursor != nil {
		defer cursor.Close(context.Background())
	}
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		var wallet model.Wallet
		err := cursor.Decode(&wallet)
		if err != nil {
			return nil, err
		}
		walletList = append(walletList, wallet)
	}
	return walletList, nil
}

func GetWalletByIdService(id primitive.ObjectID) (*model.Wallet, error) {
	filter := bson.D{{Key: "_id", Value: id}}
	var object model.Wallet

	if err := coll.WalletCollection.FindOne(context.Background(), filter).Decode(&object); err != nil {
		return nil, err
	}
	return &object, nil
}

func GetWalletByWalletNameService(name string) (*model.Wallet, error) {
	filter := bson.D{{Key: "wallet_name", Value: name}}
	var object model.Wallet

	if err := coll.WalletCollection.FindOne(context.Background(), filter).Decode(&object); err != nil {
		return nil, err
	}
	return &object, nil
}

func GetWalletByWalletIdService(id int) (*model.Wallet, error) {
	filter := bson.D{{Key: "wallet_id", Value: id}}
	var object model.Wallet

	if err := coll.WalletCollection.FindOne(context.Background(), filter).Decode(&object); err != nil {
		return nil, err
	}
	return &object, nil
}

func DeleteWalletByIdService(id primitive.ObjectID) (int, error) {
	result, err := coll.WalletCollection.DeleteOne(
		context.Background(),
		bson.D{
			{Key: "_id", Value: id},
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	return int(result.DeletedCount), nil
}

func DeleteWalletByWalletIdService(id int) int {
	result, err := coll.WalletCollection.DeleteOne(
		context.Background(),
		bson.D{
			{Key: "wallet_id", Value: id},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	return int(result.DeletedCount)
}

func UpdateWalletService(userId int, wallets model.Wallet) (*model.Wallet, error) {
	var wallet model.Wallet
	res, err := GetWalletByWalletIdService(userId)
	if err != nil {
		log.Panic()
	}
	updateWalletData := updateWalletFilter(wallets)
	updateWalletData["created_on"] = res.CreatedOn
	filter := bson.D{{Key: "wallet_id", Value: userId}}
	update := bson.D{{Key: "$set", Value: updateWalletData}}

	if err := coll.WalletCollection.FindOneAndUpdate(
		context.Background(),
		filter,
		update,
		options.FindOneAndUpdate().SetReturnDocument(1),
	).Decode(&wallet); err != nil {
		return nil, err
	}
	log.Println(wallet)
	return &wallet, nil
}

func updateWalletFilter(wallet model.Wallet) map[string]interface{} {
	var num map[string]interface{} = make(map[string]interface{})

	if wallet.WalletId != 0 {
		num["wallet_id"] = wallet.WalletId
	}
	if wallet.WalletName != "" {
		num["wallet_name"] = wallet.WalletName
	}
	if wallet.Owner != "" {
		num["owner"] = wallet.Owner
	}
	if wallet.Balance > 0 {
		num["balance"] = wallet.Balance
	}
	if wallet.Currency != "" {
		num["currency"] = wallet.Currency
	}
	if wallet.CreatedBy != "" {
		num["created_by"] = wallet.CreatedBy
	}
	date := primitive.NewDateTimeFromTime(time.Now())
	if wallet.CreatedOn != date {
		num["created_on"] = wallet.CreatedOn
	}

	if wallet.UpdatedBy != "" {
		num["updated_by"] = wallet.UpdatedBy
	}
	if wallet.UpdatedOn != date {
		num["updated_on"] = date
	}

	return num
}
