package user

import (
	"database/sql"
	"errors"
	"qk-note/shared"
)

// Service : User Service
type Service interface {
	ValidateUser(*User) error
	RegisterUser(*User) error
	UniqueUsername(user *User) error
	UniqueEmail(user *User) error
	ValidateLogin(user *Login) error
	LoginUser(user *Login) (*Claims, error)
}

// AccountService : User Account Service Struct
type AccountService struct {
	ur Repo
}

// NewAccountService : Returns User Account Service
func NewAccountService(db *sql.DB) Service {
	return &AccountService{
		ur: NewRepo(db),
	}
}

// ValidateUser : to validate User
func (as *AccountService) ValidateUser(user *User) error {
	if user.FirstName == "" {
		return errors.New("firstname required")
	} else if user.Username == "" {
		return errors.New("username required")
	} else if user.Email == "" {
		return errors.New("email required")
	} else if user.Password == "" {
		return errors.New("password required")
	}
	return nil
}

// RegisterUser : to register User
func (as *AccountService) RegisterUser(user *User) error {
	return as.ur.Register(user)
}

// UniqueUsername : to check the username already exists or not
func (as *AccountService) UniqueUsername(user *User) error {
	count, err := as.ur.CheckUsername(user)
	if err != nil {
		return nil
	} else if count != 0 {
		return errors.New("username already exists")
	}
	return nil
}

// UniqueEmail : to check the email already exists or not
func (as *AccountService) UniqueEmail(user *User) error {
	count, err := as.ur.CheckEmail(user)
	if err != nil {
		return nil
	} else if count != 0 {
		return errors.New("email already exists")
	}
	return nil
}

// ValidateLogin : to Validate Login Essentials
func (as *AccountService) ValidateLogin(user *Login) error {
	if user.Username == "" && user.Email == "" {
		return errors.New("username or email required")
	} else if user.Password == "" {
		return errors.New("password required")
	}
	return nil
}

// LoginUser : to Validate user credentials against DB
func (as *AccountService) LoginUser(login *Login) (*Claims, error) {
	user, err := as.ur.LoginUser(login)
	if err != nil {
		return nil, errors.New(shared.DatabaseError)
	} else if user == nil {
		return nil, errors.New("incorrect username or password")
	}
	var userLogged Claims
	userLogged.ID = user.ID
	userLogged.Username = user.Username
	userLogged.Email = user.Email
	return &userLogged, nil
}
