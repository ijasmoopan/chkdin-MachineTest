package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ijasmoopan/chkdin-MachineTest/models"
)

// Login is a handler function that takes in a ResponseWriter and a Request and returns nothing.
func (repo *Repo) Login(w http.ResponseWriter, r *http.Request) {

	var userLogin models.User
	var message interface{}
	json.NewDecoder(r.Body).Decode(&userLogin)
	defer r.Body.Close()

	// Checking body is empty or not
	if userLogin.Email == "" || userLogin.Password == "" {
		message = "Enter email & password"
		responseBadRequest(w, message)
		return
	}
	// checking user is valid or not
	err := repo.user.DBValidateUser(userLogin)
	if err != nil {
		message = "Incorrect email or password"
		responseBadRequest(w, message)
		return
	}
	message = "Successfully logged in"
	responseOK(w, message)
}


func (repo *Repo) PostUser(w http.ResponseWriter, r *http.Request) {

	var userSignUp models.User
	var message interface{}
	json.NewDecoder(r.Body).Decode(&userSignUp)
	defer r.Body.Close()

	// check username is already exist or not
	flag, err := repo.user.DBValidateEmail(userSignUp.Email)
	if err != nil {
		message = err
		responseBadRequest(w, message)
		return
	}
	// Adding new user
	if flag {
		err := repo.user.DBAddUser(userSignUp)
		if err != nil {
			message = err
			responseBadRequest(w, message)
			return
		}
		message = "Successfully Registered."
		responseBadRequest(w, message)
		return
	} else {
		message = "username is already exist"
		responseBadRequest(w, message)
	}
}

func (repo *Repo) GetUser(w http.ResponseWriter, r *http.Request) {
	
	email := chi.URLParam(r, "email")
	var message interface{}

	user, err := repo.user.DBGetUser(email)
	if err != nil {
		message = err
		responseBadRequest(w, message)
		return
	}
	// name := strings.Split(user.Email, "@")
	message = user
	responseOK(w, message)
}

func (repo *Repo) PatchUser(w http.ResponseWriter, r *http.Request) {

	email := chi.URLParam(r, "email")
	var patchUser models.User
	var message interface{}
	json.NewDecoder(r.Body).Decode(&patchUser)
	defer r.Body.Close()

	// edit user credentials
	if patchUser.Email == "" && patchUser.Password == "" {
		message = "Enter valid email and password"
		responseBadRequest(w, message)
		return
	}
	getUser, err := repo.user.DBGetUser(email)
	if err != nil {
		message = err
		responseBadRequest(w, message)
		return
	}
	patchUser.ID = getUser.ID
	err = repo.user.DBPatchUser(patchUser)
	if err != nil {
		message = err
		responseBadRequest(w, message)
		return
	}
	message = "User details updated"
	responseOK(w, message)
}

func (repo *Repo) DeleteUser(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	var message interface{}

	err := repo.user.DBDeleteUser(email)
	if err != nil {
		message = err
		responseBadRequest(w, message)
		return
	}
	message = "User deleted"
	responseOK(w, message)
}

