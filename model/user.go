package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id"`
	UserId      int                `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Uname       string             `json:"uname,omitempty" bson:"uname,omitempty"`
	DisplayName string             `json:"display_name,omitempty" bson:"display_name,omitempty"`
	UserRole    string             `json:"user_role,omitempty" bson:"user_role,omitempty"`
	Password    string             `json:"password,omitempty" bson:"password,omitempty"`
	Phone       []Phone            `json:"phone,omitempty" bson:"phone,omitempty"`
	Address     []Address          `json:"address,omitempty" bson:"address,omitempty"`
	FriendsList []FriendsList      `json:"friends_list,omitempty" bson:"friends_list,omitempty"`
	CreatedBy   string             `json:"created_by,omitempty" bson:"created_by,omitempty"`
	CreatedOn   primitive.DateTime `json:"created_on,omitempty" bson:"created_on,omitempty"`
	UpdatedBy   string             `json:"updated_by,omitempty" bson:"updated_by,omitempty"`
	UpdatedOn   primitive.DateTime `json:"updated_on,omitempty" bson:"updated_on,omitempty"`
}

type Book struct {
	Name      string `json:"name" bson:"name"`
	Author    string `json:"author" bson:"author"`
	PageCount int    `json:"page_count" bson:"page_count"`
}

type Phone struct {
	Number  string `json:"num,omitempty" bson:"num,omitempty"`
	Primary bool   `json:"primary" bson:"primary"`
}

type Address struct {
	Street  string `json:"street,omitempty" bson:"street,omitempty"`
	City    string `json:"city,omitempty" bson:"city,omitempty"`
	State   string `json:"state,omitempty" bson:"state,omitempty"`
	Country string `json:"country,omitempty" bson:"country,omitempty"`
}

type FriendsList struct {
	UserId          int      `json:"user_id" bson:"user_id"`
	Uname           string   `json:"uname" bson:"uname"`
	Blocked         bool     `json:"blocked" bson:"blocked"`
	BlockedForGames []string `json:"blocked_for_games" bson:"blocked_for_games"`
}
