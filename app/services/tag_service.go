package services

import (
	"notes/app/models"
	"notes/app/repositories"

	"gopkg.in/mgo.v2/bson"

	"github.com/eaciit/toolkit"

	"git.eaciitapp.com/sebar/dbflex"
)

// TagService = Service for Tag
type TagService struct {
}

// NewTagService = Create new service of Tag
func NewTagService() *TagService {
	return new(TagService)
}

// Get = Get data of Tag with dbflex filter
func (N *TagService) Get(Clause *dbflex.Filter) ([]*models.TagModel, error) {
	result := make([]*models.TagModel, 0)

	var filter *dbflex.Filter
	if Clause != nil {
		filter = Clause
	}

	param := repositories.GetByParam{
		TableName: models.NewTagModel().CollName(),
		Clause:    filter,
		Result:    &result,
	}
	err := repositories.NewRepository().GetBy(param)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetAll = Get all data of Tag
func (N *TagService) GetAll(sortKey, sortBy string, skip, limit int, filter toolkit.M) ([]*models.TagModel, int, error) {
	resultRows := make([]*models.TagModel, 0)
	var resultTotal int
	param := repositories.GetAllParam{
		TableName:   models.NewTagModel().CollName(),
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

// Insert = Insert New Tag Data
func (N *TagService) Insert(data *models.TagModel) error {
	if data.ID == "" {
		data.ID = bson.NewObjectId()
	}

	param := repositories.InsertParam{
		TableName: models.NewTagModel().CollName(),
		Data:      data,
	}

	err := repositories.NewRepository().Insert(param)

	if err != nil {
		return err
	}

	return nil
}

// Update = Update existing data
func (N *TagService) Update(data *models.TagModel) error {
	param := repositories.SaveParam{
		TableName: models.NewTagModel().CollName(),
		Data:      data,
	}

	err := repositories.NewRepository().Save(param)

	if err != nil {
		return err
	}

	return nil
}

// Delete = Delete Tag Data
func (N *TagService) Delete(clause *dbflex.Filter) error {
	param := repositories.DeleteParam{
		TableName: models.NewTagModel().CollName(),
		Clause:    clause,
	}

	err := repositories.NewRepository().Delete(param)

	if err != nil {
		return err
	}

	return nil
}

// Find = Find Tag Data
func (N *TagService) Find(clause *dbflex.Filter) (bool, []*models.TagModel, error) {
	res, err := N.Get(clause)
	if err != nil {
		return false, nil, err
	}

	if len(res) > 0 {
		return true, res, nil
	}

	return false, nil, nil
}
