package controller

import (
	"log"
	"testing"
	"time"

	collection "github.com/SageRiship/usermicroservice/connection"
	"github.com/SageRiship/usermicroservice/model"
	service "github.com/SageRiship/usermicroservice/service"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"

	"go.mongodb.org/mongo-driver/mongo"
)

var expectedUser = model.User{
	Id:          primitive.NewObjectID(),
	Uname:       "Rushikesh",
	UserId:      1,
	DisplayName: "Rushi",
	UserRole:    "ADMIN",
	Password:    "rushikesh",
	Phone:       nil,
	Address:     nil,
	FriendsList: nil,
	CreatedBy:   "Rushikesh",
	CreatedOn:   primitive.NewDateTimeFromTime(time.Now()),
	UpdatedBy:   "Rushikesh",
	UpdatedOn:   primitive.NewDateTimeFromTime(time.Now()),
}

func TestAddUser(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		collection.UserCollection = mt.Coll
		id := primitive.NewObjectID()
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		insertedUser, err := service.AddUserService(model.User{
			Id:          id,
			UserId:      12,
			Uname:       "firoj",
			DisplayName: "Firoj",
			UserRole:    "ADMIN",
			Password:    "rushikesh",
			Phone: []model.Phone{{Number: "12345",
				Primary: true}, {Number: "12345",
				Primary: true}},
			Address:     []model.Address{{Street: "Himalaya", City: "Nashik", State: "Maharashtra", Country: "India"}},
			FriendsList: nil,
			CreatedBy:   "Rishis",
			UpdatedBy:   "Rishis",
		})
		assert.Nil(t, err)
		assert.Equal(t, &model.User{
			Id:          insertedUser.Id,
			UserId:      12,
			Uname:       "firoj",
			DisplayName: "Firoj",
			UserRole:    "ADMIN",
			Password:    "rushikesh",
			Phone: []model.Phone{{Number: "12345",
				Primary: true}, {Number: "12345",
				Primary: true}},
			Address:     []model.Address{{Street: "Himalaya", City: "Nashik", State: "Maharashtra", Country: "India"}},
			FriendsList: nil,
			CreatedBy:   "Rishis",
			UpdatedBy:   "Rishis"}, insertedUser)
	})

	mt.Run("custom error duplicate", func(mt *mtest.T) {
		collection.UserCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "duplicate key error",
		}))

		insertedUser, err := service.AddUserService(model.User{})
		log.Println("Inserted User details :", insertedUser)
		assert.Nil(t, insertedUser)
		assert.NotNil(t, err)
		assert.True(t, mongo.IsDuplicateKeyError(err))
	})

	mt.Run("simple error", func(mt *mtest.T) {
		collection.UserCollection = mt.Coll
		mt.AddMockResponses(bson.D{{Key: "ok", Value: 0}})

		insertedUser, err := service.AddUserService(model.User{})

		assert.Nil(t, insertedUser)
		assert.NotNil(t, err)
	})
}

func TestGetAllUsers(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		collection.UserCollection = mt.Coll
		id1 := primitive.NewObjectID()
		id2 := primitive.NewObjectID()

		first := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: id1},
			{Key: "uname", Value: expectedUser.Uname},
			{Key: "user_id", Value: expectedUser.UserId},
			{Key: "display_name", Value: expectedUser.DisplayName},
			{Key: "user_role", Value: "ADMIN"},
			{Key: "password", Value: "rushikesh"},
			{Key: "phone", Value: expectedUser.Phone},
			{Key: "address", Value: expectedUser.Address},
			{Key: "friends_list", Value: nil},
			{Key: "created_by", Value: expectedUser.CreatedBy},
			{Key: "created_on", Value: expectedUser.CreatedOn},
			{Key: "updated_by", Value: expectedUser.UpdatedBy},
			{Key: "updated_on", Value: expectedUser.UpdatedOn},
		})
		second := mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, bson.D{
			{Key: "_id", Value: id2},
			{Key: "uname", Value: expectedUser.Uname},
			{Key: "user_id", Value: expectedUser.UserId},
			{Key: "display_name", Value: expectedUser.DisplayName},
			{Key: "user_role", Value: "ADMIN"},
			{Key: "password", Value: "rushikesh"},
			{Key: "phone", Value: expectedUser.Phone},
			{Key: "address", Value: expectedUser.Address},
			{Key: "friends_list", Value: nil},
			{Key: "created_by", Value: expectedUser.CreatedBy},
			{Key: "created_on", Value: expectedUser.CreatedOn},
			{Key: "updated_by", Value: expectedUser.UpdatedBy},
			{Key: "updated_on", Value: expectedUser.UpdatedOn},
		})
		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(first, second, killCursors)

		first_data := expectedUser
		first_data.Id = id1

		second_data := first_data
		second_data.Id = id2

		users, err := service.GetAllUsersService()

		log.Println(users)
		assert.Nil(t, err)
		assert.Equal(t, []model.User{
			first_data,
			second_data,
		}, users)
	})
}

func TestGetUserById(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		collection.UserCollection = mt.Coll
		r, _ := primitive.ObjectIDFromHex("6271f19f53d1db5e111512ac")
		expectedUser := model.User{
			Id:          r,
			Uname:       "rushi",
			UserId:      1,
			DisplayName: "rushi",
			UserRole:    "ADMIN",
			Password:    "rushikesh",
			Phone:       nil,
			Address:     nil,
			FriendsList: nil,
			CreatedBy:   "paymode",
			CreatedOn:   primitive.NewDateTimeFromTime(time.Now()),
			UpdatedBy:   "paymode",
			UpdatedOn:   primitive.NewDateTimeFromTime(time.Now()),
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: r},
			{Key: "uname", Value: expectedUser.Uname},
			{Key: "user_id", Value: expectedUser.UserId},
			{Key: "display_name", Value: expectedUser.DisplayName},
			{Key: "user_role", Value: "ADMIN"},
			{Key: "password", Value: "rushikesh"},
			{Key: "phone", Value: expectedUser.Phone},
			{Key: "address", Value: expectedUser.Address},
			{Key: "friends_list", Value: nil},
			{Key: "created_by", Value: expectedUser.CreatedBy},
			{Key: "created_on", Value: expectedUser.CreatedOn},
			{Key: "updated_by", Value: expectedUser.UpdatedBy},
			{Key: "updated_on", Value: expectedUser.UpdatedOn},
		}))
		userResponse, err := service.GetUserByIdService(r)
		log.Println(userResponse)
		assert.Nil(t, err)
		assert.Equal(t, &expectedUser, userResponse)
	})
}

func TestGetUserByUserId(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		collection.UserCollection = mt.Coll
		r, _ := primitive.ObjectIDFromHex("6271f19f53d1db5e111512ac")

		expectedUser := model.User{
			Id:          primitive.NewObjectID(),
			Uname:       "rushi",
			UserId:      1,
			DisplayName: "rushi",
			UserRole:    "ADMIN",
			Password:    "rushikesh",
			Phone:       nil,
			Address:     nil,
			FriendsList: nil,
			CreatedBy:   "paymode",
			CreatedOn:   primitive.NewDateTimeFromTime(time.Now()),
			UpdatedBy:   "paymode",
			UpdatedOn:   primitive.NewDateTimeFromTime(time.Now()),
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: r},
			{Key: "uname", Value: expectedUser.Uname},
			{Key: "user_id", Value: expectedUser.UserId},
			{Key: "display_name", Value: expectedUser.DisplayName},
			{Key: "user_role", Value: "ADMIN"},
			{Key: "password", Value: "rushikesh"},
			{Key: "phone", Value: expectedUser.Phone},
			{Key: "address", Value: expectedUser.Address},
			{Key: "friends_list", Value: nil},
			{Key: "created_by", Value: expectedUser.CreatedBy},
			{Key: "created_on", Value: expectedUser.CreatedOn},
			{Key: "updated_by", Value: expectedUser.UpdatedBy},
			{Key: "updated_on", Value: expectedUser.UpdatedOn},
		}))
		userid := expectedUser.UserId
		userResponse, err := service.GetUserByUserIdService(userid)
		log.Println(userResponse)
		assert.Nil(t, err)
		assert.NotEqual(t, &expectedUser, userResponse)
	})
}

func TestGetUserByName(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		collection.UserCollection = mt.Coll
		r, _ := primitive.ObjectIDFromHex("6271f19f53d1db5e111512ac")

		expectedUser := model.User{
			Id:          r,
			Uname:       "rushi",
			UserId:      1,
			DisplayName: "rushi",
			UserRole:    "ADMIN",
			Password:    "rushikesh",
			Phone:       nil,
			Address:     nil,
			FriendsList: nil,
			CreatedBy:   "paymode",
			CreatedOn:   primitive.NewDateTimeFromTime(time.Now()),
			UpdatedBy:   "paymode",
			UpdatedOn:   primitive.NewDateTimeFromTime(time.Now()),
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: r},
			{Key: "uname", Value: expectedUser.Uname},
			{Key: "user_id", Value: expectedUser.UserId},
			{Key: "display_name", Value: expectedUser.DisplayName},
			{Key: "user_role", Value: "ADMIN"},
			{Key: "password", Value: "rushikesh"},
			{Key: "phone", Value: expectedUser.Phone},
			{Key: "address", Value: expectedUser.Address},
			{Key: "friends_list", Value: nil},
			{Key: "created_by", Value: expectedUser.CreatedBy},
			{Key: "created_on", Value: expectedUser.CreatedOn},
			{Key: "updated_by", Value: expectedUser.UpdatedBy},
			{Key: "updated_on", Value: expectedUser.UpdatedOn},
		}))
		userName := expectedUser.Uname
		userResponse, err := service.GetUserByNameService(userName)
		log.Println(userResponse)
		assert.Nil(t, err)
		assert.Equal(t, &expectedUser, userResponse)
	})
}

func TestDeleteUserById(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		collection.UserCollection = mt.Coll
		r, _ := primitive.ObjectIDFromHex("6271f19f53d1db5e111512ac")
		expectedUser := model.User{
			Id:          r,
			Uname:       "rushi",
			UserId:      1,
			DisplayName: "rushi",
			UserRole:    "ADMIN",
			Password:    "rushikesh",
			Phone:       nil,
			Address:     nil,
			FriendsList: nil,
			CreatedBy:   "paymode",
			CreatedOn:   primitive.NewDateTimeFromTime(time.Now()),
			UpdatedBy:   "paymode",
			UpdatedOn:   primitive.NewDateTimeFromTime(time.Now()),
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: r},
			{Key: "uname", Value: expectedUser.Uname},
			{Key: "user_id", Value: expectedUser.UserId},
			{Key: "display_name", Value: expectedUser.DisplayName},
			{Key: "user_role", Value: "ADMIN"},
			{Key: "password", Value: "rushikesh"},
			{Key: "phone", Value: expectedUser.Phone},
			{Key: "address", Value: expectedUser.Address},
			{Key: "friends_list", Value: nil},
			{Key: "created_by", Value: expectedUser.CreatedBy},
			{Key: "created_on", Value: expectedUser.CreatedOn},
			{Key: "updated_by", Value: expectedUser.UpdatedBy},
			{Key: "updated_on", Value: expectedUser.UpdatedOn},
		}))
		userResponse, err := service.DeleteUserByIdService(r)
		log.Println(userResponse)
		assert.Nil(t, err)

	})
}

func TestUpdateUser(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		collection.UserCollection = mt.Coll
		r, _ := primitive.ObjectIDFromHex("6268f2a22ed1269b68417d61")
		expectedUser := model.User{
			Id:          r,
			Uname:       "rushi",
			UserId:      1,
			DisplayName: "rushi",
			UserRole:    "ADMIN",
			Password:    "rushikesh",
			Phone:       nil,
			Address:     nil,
			FriendsList: nil,
			CreatedBy:   "paymode",
			CreatedOn:   primitive.NewDateTimeFromTime(time.Now()),
			UpdatedBy:   "paymode",
			UpdatedOn:   primitive.NewDateTimeFromTime(time.Now()),
		}
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "value", Value: bson.D{
				{Key: "_id", Value: r},
				{Key: "uname", Value: expectedUser.Uname},
				{Key: "user_id", Value: expectedUser.UserId},
				{Key: "display_name", Value: expectedUser.DisplayName},
				{Key: "user_role", Value: "ADMIN"},
				{Key: "password", Value: "rushikesh"},
				{Key: "phone", Value: expectedUser.Phone},
				{Key: "address", Value: expectedUser.Address},
				{Key: "friends_list", Value: nil},
				{Key: "created_by", Value: expectedUser.CreatedBy},
				{Key: "created_on", Value: expectedUser.CreatedOn},
				{Key: "updated_by", Value: expectedUser.UpdatedBy},
				{Key: "updated_on", Value: expectedUser.UpdatedOn},
			}},
		})

		updateParams := map[string]interface{}{
			"uname": "charan",
		}
		log.Println(updateParams)

		result, err := service.UpdateUserService(expectedUser.UserId, expectedUser)

		log.Println(result)

		assert.Nil(t, err)
		assert.Equal(t, &expectedUser, result)
	})
}
