package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/SageRiship/usermicroservice/model"
	service "github.com/SageRiship/usermicroservice/service"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddWalletTransaction(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var wallet model.WalletTransaction
	_ = json.NewDecoder(request.Body).Decode(&wallet)

	wallet.Id = primitive.NewObjectID()

	result, err := service.AddWalletTransactionService(wallet)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(response).Encode(result)

}

func GetAllWalletTransaction(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var wallets []model.WalletTransaction
	wallets, err := service.GetAllWalletTransactionService()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(response).Encode(wallets)

}

func GetWalletTransactionById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	fmt.Println("Params in Get :", params)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	fmt.Println("Object Id :", id)
	wallet, err := service.GetWalletTransactionByIdService(id)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(wallet)
}

func GetWalletTransactionByOwnerName(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	fmt.Println("Params in Get :", params)
	name := params["name"]
	wallet, err := service.GetWalletTransactionByWalletOwnerService(name)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(wallet)
}

func DeleteWalletTransactionById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	count, _ := service.DeleteWalletTransactionByIdService(id)
	json.NewEncoder(response).Encode(count)
}

func DeleteWalletByWalletOwner(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	walletOwner := params["name"]

	count := service.DeleteWalletTransactionByWalletOwnerService(walletOwner)
	json.NewEncoder(response).Encode(count)
}

func UpdateWalletTransaction(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	walletOwner := params["name"]

	var wallet model.WalletTransaction
	_ = json.NewDecoder(request.Body).Decode(&wallet)

	result, _ := service.UpdateWalletTransactionService(walletOwner, wallet)

	json.NewEncoder(response).Encode(result)

}
