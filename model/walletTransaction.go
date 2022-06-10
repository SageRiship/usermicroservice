package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type WalletTransaction struct {
	Id              primitive.ObjectID `json:"_id" bson:"_id" `
	TransactionType string             `json:"transaction_type" bson:"transaction_type"`
	Amount          float64            `json:"amount" bson:"amount"`
	Currency        string             `json:"currency" bson:"currency"`
	TransactionDate primitive.DateTime `json:"transaction_date" bson:"transaction_date"`
	CreatedBy       string             `json:"created_by" bson:"created_by"`
	CreatedOn       primitive.DateTime `json:"created_on" bson:"created_on"`
	UpdatedBy       string             `json:"updated_by" bson:"updated_by"`
	UpdatedOn       primitive.DateTime `json:"updated_on" bson:"updated_on"`
	Comment         string             `json:"comment" bson:"comment"`
	GameId          string             `json:"game_id" bson:"game_id"`
	PlayNo          int                `json:"play_no" bson:"play_no"`
	RoomId          string             `json:"room_id" bson:"room_id"`
	WalletOwner     string             `json:"wallet_owner" bson:"wallet_owner"`
	TransferTo      string             `json:"transfer_to" bson:"transfer_to"`
	GamePlayId      string             `json:"game_play_id" bson:"game_play_id"`
}
