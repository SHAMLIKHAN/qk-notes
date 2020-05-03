package note

// Note : Note Struct
type Note struct {
	ID       int    `json:"id,omitempty"`
	UserID   int    `json:"user_id"`
	Heading  string `json:"heading"`
	Content  string `json:"content"`
	Status   string `json:"status,omitempty"`
	Category string `json:"category,omitempty"`
	Tags     string `json:"tags,omitempty"`
}
