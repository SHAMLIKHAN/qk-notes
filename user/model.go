package user

import "github.com/dgrijalva/jwt-go"

// User : User Struct
type User struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname,omitempty"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role,omitempty"`
	Status    string `json:"status,omitempty"`
}

// Login : User Login Struct
type Login struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Claims : Details required to identify a User
type Claims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

// JWT : Access Token
type JWT struct {
	Token string `json:"token"`
}
