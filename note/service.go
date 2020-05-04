package note

import (
	"database/sql"
	"errors"
)

// ServiceInterface : User Service
type ServiceInterface interface {
	ValidateNote(*Note) error
	CreateNote(note *Note) (*Note, error)
}

// Service : Note Service Struct
type Service struct {
	nr Repo
}

// NewService : Returns Note Service
func NewService(db *sql.DB) ServiceInterface {
	return &Service{
		nr: NewRepo(db),
	}
}

// ValidateNote : to validate Note
func (ns *Service) ValidateNote(note *Note) error {
	if note.Heading == "" {
		return errors.New("heading required")
	} else if note.Content == "" {
		return errors.New("content required")
	}
	return nil
}

// CreateNote : to create Note
func (ns *Service) CreateNote(note *Note) (*Note, error) {
	return ns.nr.Create(note)
}
