package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/SageRiship/usermicroservice/model"
	service "github.com/SageRiship/usermicroservice/service"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddWallet(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var wallet model.Wallet
	_ = json.NewDecoder(request.Body).Decode(&wallet)

	wallet.Id = primitive.NewObjectID()
	result, err := service.AddWalletService(wallet)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(response).Encode(result)

}

func GetAllWallet(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var wallets []model.Wallet
	wallets, err := service.GetAllWalletService()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(response).Encode(wallets)

}

func GetWalletById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	fmt.Println("Params in Get :", params)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	fmt.Println("Object Id :", id)
	wallet, err := service.GetWalletByIdService(id)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(wallet)
}

func GetWalletByName(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	fmt.Println("Params in Get :", params)
	name := params["name"]
	wallet, err := service.GetWalletByWalletNameService(name)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(wallet)
}

func GetWalletByUserId(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	fmt.Println("Params in Get :", params)
	walletId, _ := strconv.Atoi(params["walletid"])

	walet, err := service.GetWalletByWalletIdService(walletId)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(walet)
}

func DeleteWalletById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	count, _ := service.DeleteWalletByIdService(id)
	json.NewEncoder(response).Encode(count)
}

func DeleteWalletByWalletId(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	walletId, _ := strconv.Atoi(params["walletid"])

	count := service.DeleteWalletByWalletIdService(walletId)
	json.NewEncoder(response).Encode(count)
}

func UpdateWallet(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)

	userId, _ := strconv.Atoi(params["id"])
	var wallet model.Wallet
	_ = json.NewDecoder(request.Body).Decode(&wallet)

	result, _ := service.UpdateWalletService(userId, wallet)

	json.NewEncoder(response).Encode(result)

}
