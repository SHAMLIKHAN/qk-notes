package user

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
