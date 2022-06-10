package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Wallet struct {
	Id         primitive.ObjectID `json:"_id" bson:"_id"`
	WalletId   int                `json:"wallet_id" bson:"wallet_id"`
	WalletName string             `json:"wallet_name" bson:"wallet_name"`
	Owner      string             `json:"owner" bson:"owner"`
	Balance    float64            `json:"balance" bson:"balance"`
	Currency   string             `json:"currency" bson:"currency"`
	CreatedBy  string             `json:"created_by" bson:"created_by"`
	CreatedOn  primitive.DateTime `json:"created_on" bson:"created_on"`
	UpdatedBy  string             `json:"updated_by" bson:"updated_by"`
	UpdatedOn  primitive.DateTime `json:"updated_on" bson:"updated_on"`
}
