package note

import (
	"database/sql"
	"errors"
)

// Service : User Service
type Service interface {
	ValidateNote(*Note) error
	CreateNote(note *Note) error
}

// ServiceNote : Note Service Struct
type ServiceNote struct {
	nr Repo
}

// NewNoteService : Returns Note Service
func NewNoteService(db *sql.DB) Service {
	return &ServiceNote{
		nr: NewRepo(db),
	}
}

// ValidateNote : to validate Note
func (ns *ServiceNote) ValidateNote(note *Note) error {
	if note.Heading == "" {
		return errors.New("heading required")
	} else if note.Content == "" {
		return errors.New("content required")
	}
	return nil
}

// CreateNote : to create Note
func (ns *ServiceNote) CreateNote(note *Note) error {
	return ns.nr.Create(note)
}
