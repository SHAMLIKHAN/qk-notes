package user

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"qk-note/shared"
)

// HandlerInterface : User Handler
type HandlerInterface interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)
}

// Handler : User Handler Struct
type Handler struct {
	us ServiceInterface
}

// NewHTTPHandler : Returns User HTTP Handler
func NewHTTPHandler(db *sql.DB) HandlerInterface {
	return &Handler{
		us: NewService(db),
	}
}

// RegisterUser : to Register User
func (uh *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	log.Println("App : POST /user/register API hit!")
	var user User
	body := json.NewDecoder(r.Body)
	err := body.Decode(&user)
	if err != nil {
		log.Println("App : Error! ", err.Error())
		shared.Fail(w, 400, shared.DecodeErrorCode, shared.DecodeError)
		return
	}
	err = uh.us.ValidateUser(&user)
	if err != nil {
		log.Println("App : Error! ", err.Error())
		shared.Fail(w, 400, shared.InputDataErrorCode, err.Error())
		return
	}
	user.Role = USER
	err = uh.us.UniqueUsername(&user)
	if err != nil {
		log.Println("App : Error! ", err.Error())
		shared.Fail(w, 500, shared.DatabaseErrorCode, err.Error())
		return
	}
	err = uh.us.UniqueEmail(&user)
	if err != nil {
		log.Println("App : Error! ", err.Error())
		shared.Fail(w, 500, shared.DatabaseErrorCode, err.Error())
		return
	}
	err = uh.us.RegisterUser(&user)
	if err != nil {
		log.Println("App : Error! ", err.Error())
		shared.Fail(w, 500, shared.DatabaseErrorCode, shared.DatabaseError)
		return
	}
	log.Println("App : User registered successfully! ", user)
	shared.Send(w, 200, user)
	return
}

// LoginUser : to Login User
func (uh *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	log.Println("App : POST /user/login API hit!")
	var user Login
	body := json.NewDecoder(r.Body)
	err := body.Decode(&user)
	if err != nil {
		log.Println("App : Error! ", err.Error())
		shared.Fail(w, 400, shared.DecodeErrorCode, shared.DecodeError)
		return
	}
	err = uh.us.ValidateLogin(&user)
	if err != nil {
		log.Println("App : Error! ", err.Error())
		shared.Fail(w, 400, shared.InputDataErrorCode, err.Error())
		return
	}
	userLogged, err := uh.us.LoginUser(&user)
	if err != nil {
		log.Println("App : Error! ", err.Error())
		shared.Fail(w, 400, shared.DatabaseErrorCode, err.Error())
		return
	}
	jwtAuth := NewAuthMiddleware()
	jwt, err := jwtAuth.GenerateToken(userLogged)
	if err != nil {
		log.Println("App : Error! ", err.Error())
		shared.Fail(w, 400, shared.JWTErrorCode, shared.JWTError)
		return
	}
	shared.Send(w, 200, jwt)
}
