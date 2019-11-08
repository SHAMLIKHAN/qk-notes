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

func (c *Capsule) getNoteFromDatabase(id int) (*models.Note, error) {
	query := `SELECT * FROM note WHERE id = $1`
	result := c.DB.QueryRow(query, id)
	var note models.Note
	err := result.Scan(&note.ID, &note.UserID, &note.Title, &note.Text, &note.Status)
	if err != nil {
		return nil, err
	}
	return &note, nil
}

func (c *Capsule) updateNoteIntoDatabase(note *models.Note) error {
	query := `UPDATE note SET title = $1, text = $2 WHERE id = $3`
	_, err := c.DB.Exec(query, note.Title, note.Text, note.ID)
	if err != nil {
		return err
	}
	return nil
}

func (c *Capsule) deleteNoteFromDatabase(id int) error {
	query := `DELETE FROM note WHERE id = $1`
	_, err := c.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
