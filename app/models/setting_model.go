package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// SettingModel = Struct model of collection notes
type SettingModel struct {
	ID        bson.ObjectId `json:"_id" bson:"_id"`
	ColorTheme string      	`json:"colorTheme" bson:"colorTheme"`
	DarkTheme bool     			`json:"darkTheme" bson:"darkTheme"`
	CreatedAt time.Time     `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt" bson:"updatedAt"`
}

// NewSettingModel = Create new note model
func NewSettingModel() *SettingModel {
	n := new(SettingModel)
	n.ID = bson.NewObjectId()
	n.CreatedAt = time.Now()
	n.UpdatedAt = time.Now()
	return n
}

// CollName = Name of collection
func (N *SettingModel) CollName() string {
	return "settings"
}
