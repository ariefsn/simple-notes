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

// TagController = Controller for Tag
type TagController struct {
}

// PayloadTag = Payload for Tag
type PayloadTag struct {
	ID  bson.ObjectId `json:"_id" bson:"_id"`
	Tag string        `json:"tag" bson:"tag"`
}

// NewTagController = Create new Tag Controller
func NewTagController() *TagController {
	n := new(TagController)
	return n
}

// NewTagPayload = Create new Tag Payload
func NewTagPayload() *PayloadTag {
	n := new(PayloadTag)
	return n
}

// GetAll = Get all data for Tag
func (N *TagController) GetAll(k *knot.WebContext) {
	k.Request.Header.Set("Content-Type", "application/json")
	if core.GetMethod(k) == "GET" {
		res := core.ResponseJSON("", "", nil)
		data, err := services.NewTagService().Get(nil)
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

// GetByParam = Get all data for Tag by some params
func (N *TagController) GetByParam(k *knot.WebContext) {
	k.Request.Header.Set("Content-Type", "application/json")
	if core.GetMethod(k) == "POST" {
		payload := NewTagPayload()
		err := k.GetPayload(&payload)
		if err != nil {
			k.WriteJSON(core.ResponseJSON("nok", err.Error(), nil), http.StatusInternalServerError)
		}

		var filterAnd *dbflex.Filter
		filter := []*dbflex.Filter{}

		if payload.ID != "" {
			filter = append(filter, dbflex.Eq("_id", payload.ID))
		}
		if payload.Tag != "" {
			filter = append(filter, dbflex.Eq("tag", payload.Tag))
		}

		if len(filter) > 0 {
			filterAnd = dbflex.And(filter...)
		}

		data, err := services.NewTagService().Get(filterAnd)
		if err != nil {
			k.WriteJSON(core.ResponseJSON("nok", err.Error(), nil), http.StatusInternalServerError)
		}
		k.WriteJSON(core.ResponseJSON("ok", "success", data), http.StatusOK)
	} else {
		k.WriteJSON(core.ResponseJSONError("Unknown Method"), http.StatusInternalServerError)
	}
}

// Save = For saving data Tag
func (N *TagController) Save(k *knot.WebContext) {
	k.Request.Header.Set("Content-Type", "application/json")

	if core.GetMethod(k) == "POST" {
		payload := models.NewTagModel()

		err := k.GetPayload(&payload)
		if err != nil {
			k.WriteJSON(core.ResponseJSONError(err.Error()), http.StatusInternalServerError)
		}

		err = services.NewTagService().Insert(payload)
		if err != nil {
			k.WriteJSON(core.ResponseJSONError(err.Error()), http.StatusInternalServerError)
		}

		k.WriteJSON(core.ResponseJSONOK(nil), http.StatusCreated)
	} else {
		k.WriteJSON(core.ResponseJSONError("Unknown Method"), http.StatusInternalServerError)
	}
}

// Update = For updating data Tag
func (N *TagController) Update(k *knot.WebContext) {
	k.Request.Header.Set("Content-Type", "application/json")
	if core.GetMethod(k) == "POST" {
		payload := models.NewTagModel()

		err := k.GetPayload(&payload)
		if err != nil {
			k.WriteJSON(core.ResponseJSONError(err.Error()), http.StatusInternalServerError)
		}

		find, _, _ := services.NewTagService().Find(dbflex.Eq("_id", payload.ID))
		if !find {
			k.WriteJSON(core.ResponseJSONError("Error: Data not found!"), http.StatusOK)
		}

		err = services.NewTagService().Update(payload)
		if err != nil {
			k.WriteJSON(core.ResponseJSONError(err.Error()), http.StatusInternalServerError)
		}

		k.WriteJSON(core.ResponseJSONOK(nil), http.StatusOK)
	} else {
		k.WriteJSON(core.ResponseJSONError("Unknown Method"), http.StatusInternalServerError)
	}
}

// Delete = For delete data Tag
func (N *TagController) Delete(k *knot.WebContext) {
	k.Request.Header.Set("Content-Type", "application/json")
	if core.GetMethod(k) == "POST" {
		payload := models.NewTagModel()

		err := k.GetPayload(&payload)
		if err != nil {
			k.WriteJSON(core.ResponseJSONError(err.Error()), http.StatusInternalServerError)
		}

		find, _, _ := services.NewTagService().Find(dbflex.Eq("_id", payload.ID))
		if !find {
			k.WriteJSON(core.ResponseJSONError("Error: Data not found!"), http.StatusOK)
		}

		err = services.NewTagService().Delete(dbflex.Eq("_id", payload.ID))
		if err != nil {
			k.WriteJSON(core.ResponseJSONError(err.Error()), http.StatusInternalServerError)
		}

		k.WriteJSON(core.ResponseJSONOK(nil), http.StatusOK)
	} else {
		k.WriteJSON(core.ResponseJSONError("Unknown Method"), http.StatusInternalServerError)
	}
}
