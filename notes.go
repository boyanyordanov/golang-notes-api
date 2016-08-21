package main

import "encoding/json"

// Note: A single note
type Note struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

func (note Note) json() ([]byte, error) {
	payload, err := json.Marshal(note)
	return payload, err
}

// NoteCollection: A collection of notes
type NoteCollection struct {
	Data []Note `json:"data"`
}

func (collection NoteCollection) json() ([]byte, error) {
	payload, err := json.Marshal(collection)
	return payload, err
}

func (collection *NoteCollection) push(note Note) {
	collection.Data = append(collection.Data, note)
}
