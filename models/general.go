package models

// Note : Defines the structure of note
type Note struct {
	UserID int    `json:"user_id,omitempty"`
	Title  string `json:"title"`
	Text   string `json:"text"`
	Status int    `json:"status,omitempty"`
}
