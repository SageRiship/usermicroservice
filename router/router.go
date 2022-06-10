package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/SageRiship/usermicroservice/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	//	User Endpoints
	router.HandleFunc("/api/user", isAuthorized(controller.GetAllUsers)).Methods("GET")
	router.HandleFunc("/api/user/{id}", isAuthorized(controller.GetUserById)).Methods("GET")
	router.HandleFunc("/api/user/name/{name}", isAuthorized(controller.GetUserByName)).Methods("GET")
	router.HandleFunc("/api/user/userid/{userid}", isAuthorized(controller.GetUserByUserId)).Methods("GET")
	router.HandleFunc("/api/user", controller.AddUser).Methods("POST")
	router.HandleFunc("/api/user/{id}", isAuthorized(controller.UpdateUser)).Methods("PUT")
	router.HandleFunc("/api/user/{id}", isAuthorized(controller.DeleteUserById)).Methods("DELETE")
	router.HandleFunc("/api/user/userid/{userid}", isAuthorized(controller.DeleteUserByUserId)).Methods("DELETE")
	router.HandleFunc("/api/user", isAuthorized(controller.DeleteAllUser)).Methods("DELETE")
	router.HandleFunc("/api/user/validate", controller.ValidateUser).Methods("POST")

	// Wallet Endpoints
	router.HandleFunc("/api/wallet", isAuthorized(controller.AddWallet)).Methods("POST")
	router.HandleFunc("/api/wallet", isAuthorized(controller.GetAllWallet)).Methods("GET")
	router.HandleFunc("/api/wallet/{id}", isAuthorized(controller.GetWalletById)).Methods("GET")
	router.HandleFunc("/api/wallet/name/{name}", isAuthorized(controller.GetWalletByName)).Methods("GET")
	router.HandleFunc("/api/wallet/walletid/{walletid}", isAuthorized(controller.GetWalletByUserId)).Methods("GET")
	router.HandleFunc("/api/wallet/{id}", isAuthorized(controller.DeleteWalletById)).Methods("DELETE")
	router.HandleFunc("/api/wallet/walletid/{walletid}", isAuthorized(controller.DeleteWalletByWalletId)).Methods("DELETE")
	router.HandleFunc("/api/wallet/{id}", isAuthorized(controller.UpdateWallet)).Methods("PUT")

	// WalletTransaction Endpoints
	router.HandleFunc("/api/wallettransaction", isAuthorized(controller.AddWalletTransaction)).Methods("POST")
	router.HandleFunc("/api/wallettransaction", isAuthorized(controller.GetAllWalletTransaction)).Methods("GET")
	router.HandleFunc("/api/wallettransaction/{id}", isAuthorized(controller.GetWalletTransactionById)).Methods("GET")
	router.HandleFunc("/api/wallettransaction/name/{name}", isAuthorized(controller.GetWalletTransactionByOwnerName)).Methods("GET")
	router.HandleFunc("/api/wallettransaction/{id}", isAuthorized(controller.DeleteWalletTransactionById)).Methods("DELETE")
	router.HandleFunc("/api/wallettransaction/walletid/{name}", isAuthorized(controller.DeleteWalletByWalletOwner)).Methods("DELETE")
	router.HandleFunc("/api/wallettransaction/{name}", isAuthorized(controller.UpdateWalletTransaction)).Methods("PUT")

	return router

}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Authorization"] != nil {
			reqToken := r.Header.Get("Authorization")
			splitToken := strings.Split(reqToken, "Bearer ")
			reqToken = splitToken[1]

			type ValidationRequest struct {
				TokenString string `json:"tokenstring,omitempty"`
			}

			var validationRequest ValidationRequest
			validationRequest.TokenString = reqToken

			reqBodyBytes := new(bytes.Buffer)
			json.NewEncoder(reqBodyBytes).Encode(validationRequest)

			resp, err := http.Post("http://localhost:8200/v1/validateToken", "application/json;charset=UTF-8", reqBodyBytes) // Call to auth server for token validation.
			if err != nil {
				fmt.Printf("An error occured during REST call: %s", err.Error())
				http.Error(w, "Unable to validate token.", http.StatusInternalServerError)
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				fmt.Println("\nNon-OK HTTP status:", resp.StatusCode)
				http.Error(w, "Invalid token.", resp.StatusCode)
				return
			}

			type ValidationResponse struct {
				Name string `json:"username,omitempty"`
				Role string `json:"role,omitempty"`
			}
			var validationResponse ValidationResponse
			err = json.NewDecoder(resp.Body).Decode((&validationResponse))
			if err != nil {
				log.Printf("Unable to parse validation response %s \n", err.Error())
				http.Error(w, "Internal Error", http.StatusInternalServerError)
				return
			}

			endpoint(w, r)

		} else {
			http.Error(w, "Not Authorized", http.StatusUnauthorized)
		}
	}
}
