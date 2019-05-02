package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// CategoryModel = Struct model of collection notes
type CategoryModel struct {
	ID        bson.ObjectId `json:"_id" bson:"_id"`
	Category  string        `json:"category" bson:"category"`
	CreatedAt time.Time     `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt" bson:"updatedAt"`
}

// NewCategoryModel = Create new note model
func NewCategoryModel() *CategoryModel {
	n := new(CategoryModel)
	n.ID = bson.NewObjectId()
	n.CreatedAt = time.Now()
	n.UpdatedAt = time.Now()
	return n
}

// CollName = Name of collection
func (N *CategoryModel) CollName() string {
	return "categories"
}
