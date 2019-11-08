package core

import "qk-note/models"

func (c *Capsule) insertNoteIntoDatabase(note models.Note) error {
	query := `INSERT INTO note (user_id, title, text, status) VALUES ($1, $2, $3, $4)`
	_, err := c.DB.Exec(query, note.UserID, note.Title, note.Text, note.Status)
	if err != nil {
		return err
	}
	return nil
}

func (c *Capsule) getNotesFromDatabase(userID int) ([]models.Note, error) {
	query := `SELECT * FROM note WHERE user_id = $1`
	rows, err := c.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	var notes []models.Note
	for rows.Next() {
		var note models.Note
		err = rows.Scan(&note.ID, &note.UserID, &note.Title, &note.Text, &note.Status)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}
	return notes, nil
}
