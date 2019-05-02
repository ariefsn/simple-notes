package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// NoteModel = Struct model of collection notes
type NoteModel struct {
	ID          bson.ObjectId `json:"_id" bson:"_id"`
	Title       string        `json:"title" bson:"title"`
	Description string        `json:"description" bson:"description"`
	LinkSource  string        `json:"linkSource" bson:"linkSource"`
	Categories  []string      `json:"categories" bson:"categories"`
	Tags        []string      `json:"tags" bson:"tags"`
	CreatedAt   time.Time     `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt" bson:"updatedAt"`
}

// NewNoteModel = Create new note model
func NewNoteModel() *NoteModel {
	n := new(NoteModel)
	n.ID = bson.NewObjectId()
	n.CreatedAt = time.Now()
	n.UpdatedAt = time.Now()
	return n
}

// CollName = Name of collection
func (N *NoteModel) CollName() string {
	return "notes"
}
