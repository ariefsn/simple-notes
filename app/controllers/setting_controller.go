package controllers

import (
	"net/http"
	"notes/app/core"
	"notes/app/models"
	"notes/app/services"
	"strconv"

	"git.eaciitapp.com/sebar/dbflex"

	"gopkg.in/mgo.v2/bson"

	"git.eaciitapp.com/sebar/knot"
)

// SettingController = Controller for Setting
type SettingController struct {
}

// PayloadSetting = Payload for Setting
type PayloadSetting struct {
	ID         bson.ObjectId `json:"_id" bson:"_id"`
	ColorTheme string        `json:"colorTheme" bson:"colorTheme"`
	DarkTheme  *bool         `json:"darkTheme" bson:"darkTheme"`
}

// NewSettingController = Create new Setting Controller
func NewSettingController() *SettingController {
	n := new(SettingController)
	return n
}

// NewSettingPayload = Create new Setting Payload
func NewSettingPayload() *PayloadSetting {
	n := new(PayloadSetting)
	return n
}

// GetAll = Get all data for Setting
func (N *SettingController) GetAll(k *knot.WebContext) {
	k.Request.Header.Set("Content-Type", "application/json")
	if core.GetMethod(k) == "GET" {
		res := core.ResponseJSON("", "", nil)
		data, err := services.NewSettingService().Get(nil)
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

// GetByParam = Get all data for Setting by some params
func (N *SettingController) GetByParam(k *knot.WebContext) {
	k.Request.Header.Set("Content-Type", "application/json")
	if core.GetMethod(k) == "POST" {
		payload := NewSettingPayload()
		err := k.GetPayload(&payload)
		if err != nil {
			k.WriteJSON(core.ResponseJSON("nok", err.Error(), nil), http.StatusInternalServerError)
		}

		var filterAnd *dbflex.Filter
		filter := []*dbflex.Filter{}

		if payload.ID != "" {
			filter = append(filter, dbflex.Eq("_id", payload.ID))
		}
		if payload.ColorTheme != "" {
			filter = append(filter, dbflex.Eq("colorTheme", payload.ColorTheme))
		}
		if payload.DarkTheme != nil {
			filter = append(filter, dbflex.Eq("darkTheme", payload.DarkTheme))
		}

		if len(filter) > 0 {
			filterAnd = dbflex.And(filter...)
		}

		data, err := services.NewSettingService().Get(filterAnd)
		if err != nil {
			k.WriteJSON(core.ResponseJSON("nok", err.Error(), nil), http.StatusInternalServerError)
		}
		k.WriteJSON(core.ResponseJSON("ok", "success", data), http.StatusOK)
	} else {
		k.WriteJSON(core.ResponseJSONError("Unknown Method"), http.StatusInternalServerError)
	}
}

// Save = For saving data Setting
func (N *SettingController) Save(k *knot.WebContext) {
	k.Request.Header.Set("Content-Type", "application/json")

	if core.GetMethod(k) == "POST" {
		payload := models.NewSettingModel()

		err := k.GetPayload(&payload)
		if err != nil {
			k.WriteJSON(core.ResponseJSONError(err.Error()), http.StatusInternalServerError)
		}

		err = services.NewSettingService().Insert(payload)
		if err != nil {
			k.WriteJSON(core.ResponseJSONError(err.Error()), http.StatusInternalServerError)
		}

		k.WriteJSON(core.ResponseJSONOK(nil), http.StatusCreated)
	} else {
		k.WriteJSON(core.ResponseJSONError("Unknown Method"), http.StatusInternalServerError)
	}
}

// Update = For updating data Setting
func (N *SettingController) Update(k *knot.WebContext) {
	k.Request.Header.Set("Content-Type", "application/json")
	if core.GetMethod(k) == "POST" {
		payload := models.NewSettingModel()

		err := k.GetPayload(&payload)
		if err != nil {
			k.WriteJSON(core.ResponseJSONError(err.Error()), http.StatusInternalServerError)
		}

		find, res, _ := services.NewSettingService().Find(dbflex.Eq("_id", payload.ID))
		if !find {
			k.WriteJSON(core.ResponseJSONError("Error: Data not found!"), http.StatusOK)
		}

		// toolkit.Println(">>> = ", payload)

		res[0].ID = payload.ID
		if payload.ColorTheme != "" {
			res[0].ColorTheme = payload.ColorTheme
		}
		if strconv.FormatBool(payload.DarkTheme) != "" {
			res[0].DarkTheme = payload.DarkTheme
		}

		err = services.NewSettingService().Update(res[0])
		if err != nil {
			k.WriteJSON(core.ResponseJSONError(err.Error()), http.StatusInternalServerError)
		}

		k.WriteJSON(core.ResponseJSONOK(nil), http.StatusOK)
	} else {
		k.WriteJSON(core.ResponseJSONError("Unknown Method"), http.StatusInternalServerError)
	}
}

// Delete = For delete data Setting
func (N *SettingController) Delete(k *knot.WebContext) {
	k.Request.Header.Set("Content-Type", "application/json")
	if core.GetMethod(k) == "POST" {
		payload := models.NewSettingModel()

		err := k.GetPayload(&payload)
		if err != nil {
			k.WriteJSON(core.ResponseJSONError(err.Error()), http.StatusInternalServerError)
		}

		find, _, _ := services.NewSettingService().Find(dbflex.Eq("_id", payload.ID))
		if !find {
			k.WriteJSON(core.ResponseJSONError("Error: Data not found!"), http.StatusOK)
		}

		err = services.NewSettingService().Delete(dbflex.Eq("_id", payload.ID))
		if err != nil {
			k.WriteJSON(core.ResponseJSONError(err.Error()), http.StatusInternalServerError)
		}

		k.WriteJSON(core.ResponseJSONOK(nil), http.StatusOK)
	} else {
		k.WriteJSON(core.ResponseJSONError("Unknown Method"), http.StatusInternalServerError)
	}
}
