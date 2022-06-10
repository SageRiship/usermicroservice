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

func AddWalletTransactionService(walletTransaction model.WalletTransaction) (*model.WalletTransaction, error) {
	walletTransaction.Id = primitive.NewObjectID()
	walletTransaction.TransactionDate = primitive.NewDateTimeFromTime(time.Now())
	walletTransaction.CreatedOn = primitive.NewDateTimeFromTime(time.Now())
	walletTransaction.UpdatedOn = primitive.NewDateTimeFromTime(time.Now())
	inserted, err := coll.WalletTransactionCollection.InsertOne(context.Background(), walletTransaction)

	if err != nil {

		return nil, err
	}
	fmt.Println("Inserted 1 WalletTransaction in db with id: ", inserted.InsertedID)
	return &walletTransaction, nil

}

func GetAllWalletTransactionService() ([]model.WalletTransaction, error) {

	var walletList []model.WalletTransaction
	cursor, err := coll.WalletTransactionCollection.Find(context.Background(), bson.D{})
	if cursor != nil {
		defer cursor.Close(context.Background())
	}
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		var wallet model.WalletTransaction
		err := cursor.Decode(&wallet)
		if err != nil {
			return nil, err
		}
		walletList = append(walletList, wallet)
	}
	return walletList, nil
}

func GetWalletTransactionByIdService(id primitive.ObjectID) (*model.WalletTransaction, error) {
	filter := bson.D{{Key: "_id", Value: id}}
	var object model.WalletTransaction

	if err := coll.WalletTransactionCollection.FindOne(context.Background(), filter).Decode(&object); err != nil {
		return nil, err
	}
	return &object, nil
}

func GetWalletTransactionByWalletOwnerService(name string) ([]model.WalletTransaction, error) {
	filter := bson.D{{Key: "wallet_owner", Value: name}}
	var walletList []model.WalletTransaction
	cursor, err := coll.WalletTransactionCollection.Find(context.Background(), filter)
	if cursor != nil {
		defer cursor.Close(context.Background())
	}
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		var wallet model.WalletTransaction
		err := cursor.Decode(&wallet)
		if err != nil {
			return nil, err
		}
		walletList = append(walletList, wallet)
	}
	return walletList, nil
}

func DeleteWalletTransactionByIdService(id primitive.ObjectID) (int, error) {
	result, err := coll.WalletTransactionCollection.DeleteOne(
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

func DeleteWalletTransactionByWalletOwnerService(name string) int {
	result, err := coll.WalletTransactionCollection.DeleteOne(
		context.Background(),
		bson.D{
			{Key: "wallet_owner", Value: name},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	return int(result.DeletedCount)
}

func UpdateWalletTransactionService(walletowner string, wallets model.WalletTransaction) (*model.WalletTransaction, error) {
	var wallet model.WalletTransaction
	res, err := GetWalletTransactionByWalletOwnerService(walletowner)
	if err != nil {
		log.Panic()
	}
	updateWalletData := updateWalletTransactionFilter(wallets)
	updateWalletData["transaction_date"] = res[0].TransactionDate
	updateWalletData["created_on"] = res[0].CreatedOn
	filter := bson.D{{Key: "wallet_owner", Value: walletowner}}
	update := bson.D{{Key: "$set", Value: updateWalletData}}

	if err := coll.WalletTransactionCollection.FindOneAndUpdate(
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

func updateWalletTransactionFilter(wallet model.WalletTransaction) map[string]interface{} {
	var num map[string]interface{} = make(map[string]interface{})

	if wallet.TransactionType != "" {
		num["transaction_type"] = wallet.TransactionType
	}
	if wallet.Amount > 0 {
		num["amount"] = wallet.Amount
	}
	if wallet.Currency != "" {
		num["currency"] = wallet.Currency
	}
	if wallet.TransactionDate != primitive.NewDateTimeFromTime(time.Now()) {
		num["transaction_date"] = wallet.TransactionDate
	}
	if wallet.CreatedBy != "" {
		num["created_by"] = wallet.CreatedBy
	}
	if wallet.CreatedOn != primitive.NewDateTimeFromTime(time.Now()) {
		num["created_on"] = wallet.CreatedOn
	}
	if wallet.UpdatedBy != "" {
		num["created_by"] = wallet.UpdatedBy
	}
	if wallet.UpdatedOn != primitive.NewDateTimeFromTime(time.Now()) {
		num["updated_on"] = primitive.NewDateTimeFromTime(time.Now())
	}
	if wallet.GameId != "" {
		num["game_id"] = wallet.GameId
	}
	if wallet.PlayNo > 0 {
		num["play_no"] = wallet.PlayNo
	}
	if wallet.RoomId != "" {
		num["room_id"] = wallet.RoomId
	}
	if wallet.WalletOwner != "" {
		num["wallet_owner"] = wallet.WalletOwner
	}
	if wallet.TransferTo != "" {
		num["transfer_to"] = wallet.TransferTo
	}
	if wallet.GamePlayId != "" {
		num["game_play_id"] = wallet.GamePlayId
	}

	return num
}
