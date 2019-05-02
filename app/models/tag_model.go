package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// TagModel = Struct model of collection notes
type TagModel struct {
	ID        bson.ObjectId `json:"_id" bson:"_id"`
	Tag       string        `json:"tag" bson:"tag"`
	CreatedAt time.Time     `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt" bson:"updatedAt"`
}

// NewTagModel = Create new note model
func NewTagModel() *TagModel {
	n := new(TagModel)
	n.ID = bson.NewObjectId()
	n.CreatedAt = time.Now()
	n.UpdatedAt = time.Now()
	return n
}

// CollName = Name of collection
func (N *TagModel) CollName() string {
	return "tags"
}
