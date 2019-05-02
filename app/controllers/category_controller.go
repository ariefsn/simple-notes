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

// CategoryController = Controller for Category
type CategoryController struct {
}

// PayloadCategory = Payload for Category
type PayloadCategory struct {
	ID       bson.ObjectId `json:"_id" bson:"_id"`
	Category string        `json:"category" bson:"category"`
}

// NewCategoryController = Create new Category Controller
func NewCategoryController() *CategoryController {
	n := new(CategoryController)
	return n
}

// NewCategoryPayload = Create new Category Payload
func NewCategoryPayload() *PayloadCategory {
	n := new(PayloadCategory)
	return n
}

// GetAll = Get all data for Category
func (N *CategoryController) GetAll(k *knot.WebContext) {
	k.Request.Header.Set("Content-Type", "application/json")
	if core.GetMethod(k) == "GET" {
		res := core.ResponseJSON("", "", nil)
		data, err := services.NewCategoryService().Get(nil)
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

// GetByParam = Get all data for Category by some params
func (N *CategoryController) GetByParam(k *knot.WebContext) {
	k.Request.Header.Set("Content-Type", "application/json")
	if core.GetMethod(k) == "POST" {
		payload := NewCategoryPayload()
		err := k.GetPayload(&payload)
		if err != nil {
			k.WriteJSON(core.ResponseJSON("nok", err.Error(), nil), http.StatusInternalServerError)
		}

		var filterAnd *dbflex.Filter
		filter := []*dbflex.Filter{}

		if payload.ID != "" {
			filter = append(filter, dbflex.Eq("_id", payload.ID))
		}
		if payload.Category != "" {
			filter = append(filter, dbflex.Eq("category", payload.Category))
		}

		if len(filter) > 0 {
			filterAnd = dbflex.And(filter...)
		}

		data, err := services.NewCategoryService().Get(filterAnd)
		if err != nil {
			k.WriteJSON(core.ResponseJSON("nok", err.Error(), nil), http.StatusInternalServerError)
		}
		k.WriteJSON(core.ResponseJSON("ok", "success", data), http.StatusOK)
	} else {
		k.WriteJSON(core.ResponseJSONError("Unknown Method"), http.StatusInternalServerError)
	}
}

// Save = For saving data Category
func (N *CategoryController) Save(k *knot.WebContext) {
	k.Request.Header.Set("Content-Type", "application/json")

	if core.GetMethod(k) == "POST" {
		payload := models.NewCategoryModel()

		err := k.GetPayload(&payload)
		if err != nil {
			k.WriteJSON(core.ResponseJSONError(err.Error()), http.StatusInternalServerError)
		}

		err = services.NewCategoryService().Insert(payload)
		if err != nil {
			k.WriteJSON(core.ResponseJSONError(err.Error()), http.StatusInternalServerError)
		}

		k.WriteJSON(core.ResponseJSONOK(nil), http.StatusCreated)
	} else {
		k.WriteJSON(core.ResponseJSONError("Unknown Method"), http.StatusInternalServerError)
	}
}

// Update = For updating data Category
func (N *CategoryController) Update(k *knot.WebContext) {
	k.Request.Header.Set("Content-Type", "application/json")
	if core.GetMethod(k) == "POST" {
		payload := models.NewCategoryModel()

		err := k.GetPayload(&payload)
		if err != nil {
			k.WriteJSON(core.ResponseJSONError(err.Error()), http.StatusInternalServerError)
		}

		find, _, _ := services.NewCategoryService().Find(dbflex.Eq("_id", payload.ID))
		if !find {
			k.WriteJSON(core.ResponseJSONError("Error: Data not found!"), http.StatusOK)
		}

		err = services.NewCategoryService().Update(payload)
		if err != nil {
			k.WriteJSON(core.ResponseJSONError(err.Error()), http.StatusInternalServerError)
		}

		k.WriteJSON(core.ResponseJSONOK(nil), http.StatusOK)
	} else {
		k.WriteJSON(core.ResponseJSONError("Unknown Method"), http.StatusInternalServerError)
	}
}

// Delete = For delete data Category
func (N *CategoryController) Delete(k *knot.WebContext) {
	k.Request.Header.Set("Content-Type", "application/json")
	if core.GetMethod(k) == "POST" {
		payload := models.NewCategoryModel()

		err := k.GetPayload(&payload)
		if err != nil {
			k.WriteJSON(core.ResponseJSONError(err.Error()), http.StatusInternalServerError)
		}

		find, _, _ := services.NewCategoryService().Find(dbflex.Eq("_id", payload.ID))
		if !find {
			k.WriteJSON(core.ResponseJSONError("Error: Data not found!"), http.StatusOK)
		}

		err = services.NewCategoryService().Delete(dbflex.Eq("_id", payload.ID))
		if err != nil {
			k.WriteJSON(core.ResponseJSONError(err.Error()), http.StatusInternalServerError)
		}

		k.WriteJSON(core.ResponseJSONOK(nil), http.StatusOK)
	} else {
		k.WriteJSON(core.ResponseJSONError("Unknown Method"), http.StatusInternalServerError)
	}
}
