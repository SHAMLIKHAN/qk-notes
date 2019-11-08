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
