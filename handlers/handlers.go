package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ijasmoopan/chkdin-MachineTest/models"
)

// Login is a handler function that takes in a ResponseWriter and a Request and returns nothing.
func Login(w http.ResponseWriter, r *http.Request){

	var userLogin models.User
	json.NewDecoder(r.Body).Decode(&userLogin)
	defer r.Body.Close()

	if userLogin.UserName == "user" && userLogin.Password == "1234" {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := map[string]interface{}{
			"status": true,
			"response": "Successfully logged in",
		}
		json.NewEncoder(w).Encode(&response)
		return

	} else if userLogin.UserName == "" || userLogin.Password == "" {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]interface{}{
			"status": false,
			"response": "Enter username & password",
		}
		json.NewEncoder(w).Encode(&response)
		return

	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]interface{}{
			"status": false,
			"response": "Incorrect username or password",
		}
		json.NewEncoder(w).Encode(&response)
	}
}

// func Home(w http.ResponseWriter, r *http.Request){

// }

func GetUser(w http.ResponseWriter, r *http.Request){
	user := chi.URLParam(r, "username")

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"status": true,
		"response": fmt.Sprint("Hi ", user),
	}
	json.NewEncoder(w).Encode(&response)
}

func PostUser(w http.ResponseWriter, r *http.Request){

	var userSignUp models.User
	json.NewDecoder(r.Body).Decode(&userSignUp)
	defer r.Body.Close()

	// check username is already exist or not
	
	if userSignUp.UserName != "" && userSignUp.Password != "" {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := map[string]interface{}{
			"status": true,
			"response": "Successfully Registered.",
		}
		json.NewEncoder(w).Encode(&response)
	}

}

func PatchUser(w http.ResponseWriter, r *http.Request){
	var patchUser models.User
	json.NewDecoder(r.Body).Decode(&patchUser)

	// edit user credentials

}

func DeletUser(w http.ResponseWriter, r *http.Request){
	// user := chi.URLParam(r, "username")
	
	
	// delete user
}


