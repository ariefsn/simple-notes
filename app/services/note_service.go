package services

import (
	"notes/app/models"
	"notes/app/repositories"

	"gopkg.in/mgo.v2/bson"

	"github.com/eaciit/toolkit"

	"git.eaciitapp.com/sebar/dbflex"
)

// NoteService = Service for Note
type NoteService struct {
}

// NewNoteService = Create new service of note
func NewNoteService() *NoteService {
	return new(NoteService)
}

// Get = Get data of notes with dbflex filter
func (N *NoteService) Get(Clause *dbflex.Filter) ([]*models.NoteModel, error) {
	result := make([]*models.NoteModel, 0)

	var filter *dbflex.Filter
	if Clause != nil {
		filter = Clause
	}

	param := repositories.GetByParam{
		TableName: models.NewNoteModel().CollName(),
		Clause:    filter,
		Result:    &result,
	}
	err := repositories.NewRepository().GetBy(param)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetAll = Get all data of notes
func (N *NoteService) GetAll(sortKey, sortBy string, skip, limit int, filter toolkit.M) ([]*models.NoteModel, int, error) {
	resultRows := make([]*models.NoteModel, 0)
	var resultTotal int
	param := repositories.GetAllParam{
		TableName:   models.NewNoteModel().CollName(),
		Filter:      filter,
		Skip:        skip,
		Take:        limit,
		SortKey:     sortKey,
		SortOrder:   sortBy,
		ResultRows:  resultRows,
		ResultTotal: &resultTotal,
	}

	err := repositories.NewRepository().GetAll(param)
	if err != nil {
		return nil, 0, err
	}

	return resultRows, resultTotal, nil
}

// Insert = Insert New Note Data
func (N *NoteService) Insert(data *models.NoteModel) error {
	if data.ID == "" {
		data.ID = bson.NewObjectId()
	}

	param := repositories.InsertParam{
		TableName: models.NewNoteModel().CollName(),
		Data:      data,
	}

	err := repositories.NewRepository().Insert(param)

	if err != nil {
		return err
	}

	return nil
}

// Update = Update existing data
func (N *NoteService) Update(data *models.NoteModel) error {
	param := repositories.SaveParam{
		TableName: models.NewNoteModel().CollName(),
		Data:      data,
	}

	err := repositories.NewRepository().Save(param)

	if err != nil {
		return err
	}

	return nil
}

// Delete = Delete Note Data
func (N *NoteService) Delete(clause *dbflex.Filter) error {
	param := repositories.DeleteParam{
		TableName: models.NewNoteModel().CollName(),
		Clause:    clause,
	}

	err := repositories.NewRepository().Delete(param)

	if err != nil {
		return err
	}

	return nil
}

// Find = Find Note Data
func (N *NoteService) Find(clause *dbflex.Filter) (bool, []*models.NoteModel, error) {
	res, err := N.Get(clause)
	if err != nil {
		return false, nil, err
	}

	if len(res) > 0 {
		return true, res, nil
	}

	return false, nil, nil
}
