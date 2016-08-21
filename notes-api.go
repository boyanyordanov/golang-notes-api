package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var fileLog = FileLogger{path: "logs/notes-api-log.txt"}
var notes NoteCollection
var lastNoteId = 1

func NotePostHandler(w http.ResponseWriter, req *http.Request) {
	go fileLog.Info("Received request: POST /notes")

	decoder := json.NewDecoder(req.Body)
	var newNote Note

	// Read the data for the new note and decode it to the correct struct
	err := decoder.Decode(&newNote)

	if err != nil {
		fileLog.Error(err.Error())
	}

	// Update the id of the note
	newNote.ID = lastNoteId
	lastNoteId++

	// Push the note in the store
	notes.push(newNote)

	// Return the new note as response
	WriteJson(w, newNote, http.StatusCreated, fileLog)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		go fileLog.Info("Demo of info log")
		go fileLog.Error("Demo of error log")
		w.Write([]byte("Welcome to the notes API"))
	})

	router.HandleFunc("/notes", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "GET":
			go fileLog.Info("Received request: GET /notes")
			WriteJson(w, notes, http.StatusOK, fileLog)
		case "POST":
			NotePostHandler(w, req)
		}
	})

	log.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
