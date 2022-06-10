package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/SageRiship/usermicroservice/model"

	service "github.com/SageRiship/usermicroservice/service"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var user model.User
	_ = json.NewDecoder(request.Body).Decode(&user)

	/*
		if we want to generate id by GO write this and
			[ Id          primitive.ObjectID  `json:"_id" bson:"_id"` ]..in struct
	*/
	user.Id = primitive.NewObjectID()
	//collection = client.Database(dbname).Collection(colname)
	//ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//result, err := UserCollection.InsertOne(context.Background(), user)
	result, errs := service.GetUserByNameService(user.Uname)

	if errs != nil && result == nil {
		result, err := service.AddUserService(user)
		if err != nil {
			http.Error(response, "User Creation Failed", http.StatusUnprocessableEntity)
			return
		}
		wallet := model.Wallet{
			Id:         primitive.NewObjectID(),
			WalletId:   101,
			WalletName: user.Uname,
			Owner:      user.Uname,
			Balance:    0,
			Currency:   "USD",
			CreatedBy:  user.Uname,
			CreatedOn:  user.CreatedOn,
			UpdatedBy:  user.Uname,
			UpdatedOn:  user.UpdatedOn,
		}
		if err == nil {
			service.AddWalletService(wallet)
		}
		json.NewEncoder(response).Encode(result)
	} else {
		http.Error(response, "User already present", http.StatusConflict)
	}

}

func ValidateUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var validate model.Validation
	_ = json.NewDecoder(request.Body).Decode(&validate)

	result, err := service.ValidateUserService(validate.Uname, validate.Password)
	if err != nil {
		http.Error(response, "Unauthorised User", http.StatusUnauthorized)
		return
	}
	if result == nil {
		http.Error(response, "Unauthorised User", http.StatusUnauthorized)
		return
	}
	json.NewEncoder(response).Encode(result)
}

func GetAllUsers(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var users []model.User
	//collection = client.Database(dbname).Collection(colname)
	//ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	//cursor, err := collection.UserCollection.Find(context.Background(), bson.M{})
	users, err := service.GetAllUsersService()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(response).Encode(users)

}

func GetUserById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	params := mux.Vars(request)
	fmt.Println("Params in Get :", params)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	fmt.Println("Object Id :", id)
	user, err := service.GetUserByIdService(id)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(user)
}

func GetUserByName(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	fmt.Println("Params in Get :", params)
	name := params["name"]
	user, err := service.GetUserByNameService(name)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(user)
}

func GetUserByUserId(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	fmt.Println("Params in Get :", params)
	userId, _ := strconv.Atoi(params["userid"])

	user, err := service.GetUserByUserIdService(userId)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(user)
}

func DeleteAllUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := service.DeleteAllUser()
	json.NewEncoder(response).Encode(count)
}

func DeleteUserById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	count, _ := service.DeleteUserByIdService(id)
	json.NewEncoder(response).Encode(count)
}

func DeleteUserByUserId(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	userId, _ := strconv.Atoi(params["userid"])

	count := service.DeleteUserByUserIdService(userId)
	json.NewEncoder(response).Encode(count)
}

func UpdateUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)

	userId, _ := strconv.Atoi(params["id"])
	var user model.User
	_ = json.NewDecoder(request.Body).Decode(&user)

	result, _ := service.UpdateUserService(userId, user)

	json.NewEncoder(response).Encode(result)

}
