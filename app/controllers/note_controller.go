package controllers

import (
	"net/http"
	"notes/app/core"
	"notes/app/models"
	"notes/app/services"

	"git.eaciitapp.com/sebar/dbflex"

	"gopkg.in/mgo.v2/bson"

	"git.eaciitapp.com/sebar/knot"
)

// NoteController = Controller for note
type NoteController struct {
}

// PayloadNote = Payload for Note
type PayloadNote struct {
	ID          bson.ObjectId `json:"_id" bson:"_id"`
	Title       string        `json:"title" bson:"title"`
	Description string        `json:"description" bson:"description"`
	Tags        []interface{} `json:"tags" bson:"tags"`
	Categories  []interface{} `json:"categories" bson:"categories"`
}

// NewNoteController = Create new Note Controller
func NewNoteController() *NoteController {
	n := new(NoteController)
	return n
}

// NewNotePayload = Create new Note Payload
func NewNotePayload() *PayloadNote {
	n := new(PayloadNote)
	return n
}

// GetAll = Get all data for notes
func (N *NoteController) GetAll(k *knot.WebContext) {
	k.Request.Header.Set("Content-Type", "application/json")
	if core.GetMethod(k) == "GET" {
		res := core.ResponseJSON("", "", nil)
		data, err := services.NewNoteService().Get(nil)
		if err != nil {
			res.Set("status", "nok")
			res.Set("message", err.Error())
			k.WriteJSON(res, http.StatusInternalServerError)
		}
		res.Set("status", "ok")
		res.Set("message", "success")
		res.Set("data", data)
		k.WriteJSON(core.ResponseJSONOK(res), http.StatusOK)
	} else {
		k.WriteJSON(core.ResponseJSONError("Unknown Method"), http.StatusInternalServerError)
	}
}

// GetByParam = Get all data for notes by some params
func (N *NoteController) GetByParam(k *knot.WebContext) {
	k.Request.Header.Set("Content-Type", "application/json")
	if core.GetMethod(k) == "POST" {
		payload := NewNotePayload()
		err := k.GetPayload(&payload)
		if err != nil {
			k.WriteJSON(core.ResponseJSON("nok", err.Error(), nil), http.StatusInternalServerError)
		}

		var filterAnd *dbflex.Filter
		filter := []*dbflex.Filter{}

		if payload.ID != "" {
			filter = append(filter, dbflex.Eq("_id", payload.ID))
		}
		if payload.Title != "" {
			filter = append(filter, dbflex.Eq("title", payload.Title))
		}
		if payload.Description != "" {
			filter = append(filter, dbflex.Eq("description", payload.Description))
		}
		if payload.Tags != nil {
			filter = append(filter, dbflex.In("tags", payload.Tags...))
		}
		if payload.Categories != nil {
			filter = append(filter, dbflex.In("categories", payload.Categories...))
		}

		if len(filter) > 0 {
			filterAnd = dbflex.And(filter...)
		}

		data, err := services.NewNoteService().Get(filterAnd)
		if err != nil {
			k.WriteJSON(core.ResponseJSON("nok", err.Error(), nil), http.StatusInternalServerError)
		}
		k.WriteJSON(core.ResponseJSON("ok", "success", data), http.StatusOK)
	} else {
		k.WriteJSON(core.ResponseJSONError("Unknown Method"), http.StatusInternalServerError)
	}
}

// Save = For saving data note
func (N *NoteController) Save(k *knot.WebContext) {
	k.Request.Header.Set("Content-Type", "application/json")

	if core.GetMethod(k) == "POST" {
		payload := models.NewNoteModel()

		err := k.GetPayload(&payload)
		if err != nil {
			k.WriteJSON(core.ResponseJSONError(err.Error()), http.StatusInternalServerError)
		}

		err = services.NewNoteService().Insert(payload)
		if err != nil {
			k.WriteJSON(core.ResponseJSONError(err.Error()), http.StatusInternalServerError)
		}

		k.WriteJSON(core.ResponseJSONOK(nil), http.StatusCreated)
	} else {
		k.WriteJSON(core.ResponseJSONError("Unknown Method"), http.StatusInternalServerError)
	}
}

// Update = For updating data note
func (N *NoteController) Update(k *knot.WebContext) {
	k.Request.Header.Set("Content-Type", "application/json")
	if core.GetMethod(k) == "POST" {
		payload := models.NewNoteModel()

		err := k.GetPayload(&payload)
		if err != nil {
			k.WriteJSON(core.ResponseJSONError(err.Error()), http.StatusInternalServerError)
		}

		find, _, _ := services.NewNoteService().Find(dbflex.Eq("_id", payload.ID))
		if !find {
			k.WriteJSON(core.ResponseJSONError("Error: Data not found!"), http.StatusOK)
		}

		err = services.NewNoteService().Update(payload)
		if err != nil {
			k.WriteJSON(core.ResponseJSONError(err.Error()), http.StatusInternalServerError)
		}

		k.WriteJSON(core.ResponseJSONOK(nil), http.StatusOK)
	} else {
		k.WriteJSON(core.ResponseJSONError("Unknown Method"), http.StatusInternalServerError)
	}
}

// Delete = For delete data note
func (N *NoteController) Delete(k *knot.WebContext) {
	k.Request.Header.Set("Content-Type", "application/json")
	if core.GetMethod(k) == "POST" {
		payload := models.NewNoteModel()

		err := k.GetPayload(&payload)
		if err != nil {
			k.WriteJSON(core.ResponseJSONError(err.Error()), http.StatusInternalServerError)
		}

		find, _, _ := services.NewNoteService().Find(dbflex.Eq("_id", payload.ID))
		if !find {
			k.WriteJSON(core.ResponseJSONError("Error: Data not found!"), http.StatusOK)
		}

		err = services.NewNoteService().Delete(dbflex.Eq("_id", payload.ID))
		if err != nil {
			k.WriteJSON(core.ResponseJSONError(err.Error()), http.StatusInternalServerError)
		}

		k.WriteJSON(core.ResponseJSONOK(nil), http.StatusOK)
	} else {
		k.WriteJSON(core.ResponseJSONError("Unknown Method"), http.StatusInternalServerError)
	}
}
