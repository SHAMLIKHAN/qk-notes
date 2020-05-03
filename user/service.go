package user

import (
	"database/sql"
	"errors"
	"qk-note/shared"
)

// ServiceInterface : User Service
type ServiceInterface interface {
	ValidateUser(*User) error
	RegisterUser(*User) error
	UniqueUsername(user *User) error
	UniqueEmail(user *User) error
	ValidateLogin(user *Login) error
	LoginUser(user *Login) (*Claims, error)
}

// Service : User Service Struct
type Service struct {
	ur Repo
}

// NewService : Returns User Service
func NewService(db *sql.DB) ServiceInterface {
	return &Service{
		ur: NewRepo(db),
	}
}

// ValidateUser : to validate User
func (us *Service) ValidateUser(user *User) error {
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
func (us *Service) RegisterUser(user *User) error {
	return us.ur.Register(user)
}

// UniqueUsername : to check the username already exists or not
func (us *Service) UniqueUsername(user *User) error {
	count, err := us.ur.CheckUsername(user)
	if err != nil {
		return nil
	} else if count != 0 {
		return errors.New("username already exists")
	}
	return nil
}

// UniqueEmail : to check the email already exists or not
func (us *Service) UniqueEmail(user *User) error {
	count, err := us.ur.CheckEmail(user)
	if err != nil {
		return nil
	} else if count != 0 {
		return errors.New("email already exists")
	}
	return nil
}

// ValidateLogin : to Validate Login Essentials
func (us *Service) ValidateLogin(user *Login) error {
	if user.Username == "" && user.Email == "" {
		return errors.New("username or email required")
	} else if user.Password == "" {
		return errors.New("password required")
	}
	return nil
}

// LoginUser : to Validate user credentials against DB
func (us *Service) LoginUser(login *Login) (*Claims, error) {
	user, err := us.ur.LoginUser(login)
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
