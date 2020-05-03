package user

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"qk-note/shared"
)

// Handler : User Handler
type Handler interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
}

// AccountHandler : User Account Handler Struct
type AccountHandler struct {
	as Service
}

// NewHTTPHandler : Returns User HTTP Handler
func NewHTTPHandler(db *sql.DB) Handler {
	return &AccountHandler{
		as: NewAccountService(db),
	}
}

// RegisterUser : to Register User
func (ah *AccountHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	log.Println("App : POST /register API hit!")
	var user User
	body := json.NewDecoder(r.Body)
	err := body.Decode(&user)
	if err != nil {
		log.Println("App : Error! ", err.Error())
		shared.Fail(w, 400, shared.DecodeErrorCode, shared.DecodeError)
		return
	}
	err = ah.as.ValidateUser(&user)
	if err != nil {
		log.Println("App : Error! ", err.Error())
		shared.Fail(w, 400, shared.InputDataErrorCode, err.Error())
		return
	}
	user.Role = USER
	err = ah.as.UniqueUsername(&user)
	if err != nil {
		log.Println("App : Error! ", err.Error())
		shared.Fail(w, 500, shared.DatabaseErrorCode, err.Error())
		return
	}
	err = ah.as.UniqueEmail(&user)
	if err != nil {
		log.Println("App : Error! ", err.Error())
		shared.Fail(w, 500, shared.DatabaseErrorCode, err.Error())
		return
	}
	err = ah.as.RegisterUser(&user)
	if err != nil {
		log.Println("App : Error! ", err.Error())
		shared.Fail(w, 500, shared.DatabaseErrorCode, shared.DatabaseError)
		return
	}
	log.Println("App : User registered successfully! ", user)
	shared.Send(w, 200, user)
	return
}
