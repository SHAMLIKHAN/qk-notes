package note

import (
	"database/sql"
	"errors"
)

// ServiceInterface : User Service
type ServiceInterface interface {
	ValidateNote(*Note) error
	CreateNote(note *Note) (*Note, error)
	GetNotes(note *Note) ([]Note, error)
	FilterNotes(notes []Note, conditions *Note) []Note
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

// GetNotes : to fetch all Notes
func (ns *Service) GetNotes(note *Note) ([]Note, error) {
	notes, err := ns.nr.GetAll(note)
	if err != nil {
		return nil, err
	}
	notes = ns.FilterNotes(notes, note)
	return notes, nil
}

// FilterNotes : to filter notes based on conditions
func (ns *Service) FilterNotes(notes []Note, conditions *Note) []Note {
	// based on status
	var result1 []Note
	if conditions.Status != "" {
		for _, note := range notes {
			if conditions.Status == note.Status {
				result1 = append(result1, note)
			}
		}
		notes = result1
	}
	// based on category
	var result2 []Note
	if conditions.Category != "" {
		for _, note := range notes {
			if conditions.Category == note.Category {
				result2 = append(result2, note)
			}
		}
		notes = result2
	}
	return notes
}
