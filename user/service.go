package user

import (
	"database/sql"
	"errors"
)

// Service : User Service
type Service interface {
	ValidateUser(*User) error
	RegisterUser(*User) error
	UniqueUsername(user *User) error
	UniqueEmail(user *User) error
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
