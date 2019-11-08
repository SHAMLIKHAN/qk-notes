package core

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"qk-note/consts"
	"qk-note/models"
	"qk-note/shared"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateNote : to add a new note
func (c *Capsule) CreateNote(w http.ResponseWriter, r *http.Request) {
	var note models.Note
	body := json.NewDecoder(r.Body)
	err := body.Decode(&note)
	if err != nil {
		log.Println("App : Error! ", err.Error())
		shared.Fail(w, 400, consts.DecodeErrorCode, consts.DecodeError)
		return
	}
	note.UserID = 1
	note.Status = consts.Active
	err = c.insertNoteIntoDatabase(note)
	if err != nil {
		log.Println("App : Error! ", err.Error())
		shared.Fail(w, 500, consts.DatabaseErrorCode, consts.DatabaseError)
		return
	}
	log.Println("App : POST /note API called!")
	log.Println("App : Note inserted to database! ", note)
	shared.Send(w, 200, note)
}

// GetNotes : to return all notes
func (c *Capsule) GetNotes(w http.ResponseWriter, r *http.Request) {
	userID := 1
	notes, err := c.getNotesFromDatabase(userID)
	if err != nil {
		log.Println("App : Error! ", err.Error())
		shared.Fail(w, 500, consts.DatabaseErrorCode, consts.DatabaseError)
		return
	}
	log.Println("App : GET /note API called!")
	log.Println("App : Notes are retrieved from database! ", notes)
	shared.Send(w, 200, notes)
}

// GetNote : to get a particular note
func (c *Capsule) GetNote(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	noteID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("App: Error! ", err.Error())
		shared.Fail(w, 400, consts.InputDataErrorCode, consts.InputDataError)
		return
	}
	note, err := c.getNoteFromDatabase(noteID)
	if err != nil {
		log.Println("App : Error! ", err.Error())
		shared.Fail(w, 500, consts.DatabaseErrorCode, consts.DatabaseError)
		return
	}
	log.Println("App : GET /note/{id} API called!")
	log.Println("App : Note retrieved from database! ", note)
	shared.Send(w, 200, note)
}

// EditNote : to update a particular note
func (c *Capsule) EditNote(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	noteID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("App: Error! ", err.Error())
		shared.Fail(w, 400, consts.InputDataErrorCode, consts.InputDataError)
		return
	}
	existingNote, err := c.getNoteFromDatabase(noteID)
	if err != nil {
		log.Println("App : Error! ", err.Error())
		shared.Fail(w, 500, consts.DatabaseErrorCode, consts.DatabaseError)
		return
	}
	var editedNote models.Note
	body := json.NewDecoder(r.Body)
	err = body.Decode(&editedNote)
	if err != nil {
		log.Println("App : Error! ", err.Error())
		shared.Fail(w, 400, consts.DecodeErrorCode, consts.DecodeError)
		return
	}
	updatedNote := prepareNewNote(existingNote, &editedNote)
	err = c.updateNoteIntoDatabase(updatedNote)
	if err != nil {
		log.Println("App : Error! ", err.Error())
		shared.Fail(w, 400, consts.DatabaseErrorCode, consts.DatabaseError)
		return
	}
	log.Println("App : PATCH /note/{id} API called!")
	log.Println("App : Note updated into database! Note ID : ", noteID)
	shared.Send(w, 200, updatedNote)
}

// DeleteNote : to delete a particular note
func (c *Capsule) DeleteNote(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	noteID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("App: Error! ", err.Error())
		shared.Fail(w, 400, consts.InputDataErrorCode, consts.InputDataError)
		return
	}
	err = c.deleteNoteFromDatabase(noteID)
	if err != nil {
		log.Println("App : Error! ", err.Error())
		shared.Fail(w, 500, consts.DatabaseErrorCode, consts.DatabaseError)
		return
	}
	log.Println("App : DELETE /note/{id} API called!")
	log.Println("App : Note deleted from database! Note ID : ", noteID)
	response := `Success! Note [` + id + `] Deleted`
	shared.Send(w, 200, response)
}

func prepareNewNote(existingNote, editedNote *models.Note) *models.Note {
	var note models.Note
	note.ID = existingNote.ID
	note.UserID = existingNote.UserID
	note.Title = editedNote.Title
	note.Text = editedNote.Text
	note.Status = existingNote.Status
	return &note
}
