package models

// Note : Defines the structure of note
type Note struct {
	ID     int    `json:"id,omitempty"`
	UserID int    `json:"user_id,omitempty"`
	Title  string `json:"title"`
	Text   string `json:"text"`
	Status int    `json:"status,omitempty"`
}
