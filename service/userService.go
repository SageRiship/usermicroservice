package service

import (
	"context"
	"fmt"
	"log"
	"time"

	coll "github.com/SageRiship/usermicroservice/connection"
	enc "github.com/SageRiship/usermicroservice/encryption"
	"github.com/SageRiship/usermicroservice/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddUserService(user model.User) (*model.User, error) {

	/*
		if we want to generate id by GO write this and
			[ Id          primitive.ObjectID  `json:"_id" bson:"_id"` ]..in struct
	*/
	user.Id = primitive.NewObjectID()
	encPassword := enc.Encrypt(user.Password)
	user.Password = encPassword
	user.CreatedOn = primitive.NewDateTimeFromTime(time.Now())
	user.UpdatedOn = primitive.NewDateTimeFromTime(time.Now())
	inserted, err := coll.UserCollection.InsertOne(context.Background(), user)

	if err != nil {

		return nil, err
	}
	fmt.Println("Inserted 1 User in db with id: ", inserted.InsertedID)

	return &user, nil

}

func GetAllUsersService() ([]model.User, error) {

	var users []model.User
	//collection = client.Database(dbname).Collection(colname)
	//ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := coll.UserCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	//defer cursor.Close(ctx)
	for cursor.Next(context.Background()) {
		var user model.User
		cursor.Decode(&user)
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return users, nil

}

func GetUserByIdService(id primitive.ObjectID) (*model.User, error) {
	filter := bson.D{{Key: "_id", Value: id}}
	var object model.User

	if err := coll.UserCollection.FindOne(context.Background(), filter).Decode(&object); err != nil {
		return nil, err
	}
	return &object, nil
}

func GetUserByNameService(name string) (*model.User, error) {
	filter := bson.D{{Key: "uname", Value: name}}
	var object model.User

	if err := coll.UserCollection.FindOne(context.Background(), filter).Decode(&object); err != nil {
		return nil, err
	}
	return &object, nil
}

func GetUserByUserIdService(id int) (*model.User, error) {
	filter := bson.D{{Key: "user_id", Value: id}}
	var object model.User

	if err := coll.UserCollection.FindOne(context.Background(), filter).Decode(&object); err != nil {
		return nil, err
	}
	return &object, nil
}

func DeleteAllUser() int64 {

	deleteResult, err := coll.UserCollection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Number of User delete: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

func DeleteUserByIdService(id primitive.ObjectID) (int, error) {
	result, err := coll.UserCollection.DeleteOne(
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

func DeleteUserByUserIdService(id int) int {
	result, err := coll.UserCollection.DeleteOne(
		context.Background(),
		bson.D{
			{Key: "user_id", Value: id},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	return int(result.DeletedCount)
}

func UpdateUserService(userId int, users model.User) (*model.User, error) {
	var user model.User
	//objectId, err := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	return nil, err
	// }

	updateUserData := updateFilter(users)
	//	updateLeaderboardData["updated_on"] = primitive.Timestamp{T: uint32(time.Now().Unix())}
	filter := bson.D{{Key: "user_id", Value: userId}}
	update := bson.D{{Key: "$set", Value: updateUserData}}

	if err := coll.UserCollection.FindOneAndUpdate(
		context.Background(),
		filter,
		update,
		options.FindOneAndUpdate().SetReturnDocument(1),
	).Decode(&user); err != nil {
		return nil, err
	}
	log.Println(user)
	return &user, nil
}

func updateFilter(user model.User) map[string]interface{} {
	var num map[string]interface{} = make(map[string]interface{})

	num["userid"] = user.UserId
	num["uname"] = user.Uname
	num["display_name"] = user.DisplayName
	num["user_role"] = user.UserRole
	num["password"] = user.Password
	num["phone"] = user.Phone
	num["address"] = user.Address
	num["friends_list"] = user.FriendsList
	num["created_by"] = user.CreatedBy
	num["updated_by"] = user.UpdatedBy

	return num
}

func ValidateUserService(username, password string) (*model.User, error) {
	log.Println("User Request: ", username)

	res, err := GetUserByNameService(username)
	if err != nil {
		log.Println("User does not exit in our database:", username)
		return nil, err
	}
	log.Println("user details:", res)
	if enc.Decrypt(res.Password) != password {
		return nil, err
	}

	return res, nil
}
