package note

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"encore.dev/rlog"
)

// Type that represents a note.
type Note struct {
	ID       string `json:"id"`
	Text     string `json:"text"`
	CoverURL string `json:"cover_url"`
}

const notesFile = "/note/db/notes.json"

var (
	mu    sync.Mutex
	notes = make(map[string]*Note)
)

func loadNotes() error {
	rlog.Info("Loading notes from file")

	pwd, _ := os.Getwd()
	file, err := os.Open(filepath.Join(pwd, notesFile))
	if err != nil {
		if os.IsNotExist(err) {
			error := fmt.Sprintf("file %v not found", notesFile)
			return errors.New(error)
		}
		return err
	}
	defer file.Close()

	bytes, err := os.ReadFile(filepath.Join(pwd, notesFile))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	if err := json.Unmarshal(bytes, &notes); err != nil {
		return err
	}

	return nil
}

func saveToFile() error {
	bytes, err := json.Marshal(notes)
	if err != nil {
		return err
	}

	pwd, _ := os.Getwd()
	file, err := os.Open(filepath.Join(pwd, notesFile))
	if err != nil {
		if os.IsNotExist(err) {
			error := fmt.Sprintf("file %v not found", notesFile)
			return errors.New(error)
		}
		return err
	}
	defer file.Close()
	if err := os.WriteFile(filepath.Join(pwd, notesFile), bytes, 0644); err != nil {
		return err
	}

	return nil
}

//encore:api public method=POST path=/note
func SaveNote(ctx context.Context, note *Note) (*Note, error) {

	mu.Lock()
	defer mu.Unlock()

	notes[note.ID] = note

	// If there was an error saving to the file, then we return that error.
	if err := saveToFile(); err != nil {
		return nil, err
	}

	// Otherwise, we return the note to indicate that the save was successful.
	return note, nil
}

//encore:api public method=GET path=/note/:id
func GetNote(ctx context.Context, id string) (*Note, error) {
	rlog.Debug("Inside GetNote()", "id", id)

	mu.Lock()
	defer mu.Unlock()

	if err := loadNotes(); err != nil {
		rlog.Error("Error", "err", err)
		return nil, err
	}

	// We use the note ID to query the map for the note's text and cover URL.
	note, exists := notes[id]
	if !exists {
		return nil, errors.New("Note not found")
	}

	// Otherwise, we return the note.
	return note, nil
}
