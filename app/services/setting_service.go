package services

import (
	"notes/app/models"
	"notes/app/repositories"

	"gopkg.in/mgo.v2/bson"

	"github.com/eaciit/toolkit"

	"git.eaciitapp.com/sebar/dbflex"
)

// SettingService = Service for Setting
type SettingService struct {
}

// NewSettingService = Create new service of Setting
func NewSettingService() *SettingService {
	return new(SettingService)
}

// Get = Get data of Setting with dbflex filter
func (N *SettingService) Get(Clause *dbflex.Filter) ([]*models.SettingModel, error) {
	result := make([]*models.SettingModel, 0)

	var filter *dbflex.Filter
	if Clause != nil {
		filter = Clause
	}

	param := repositories.GetByParam{
		TableName: models.NewSettingModel().CollName(),
		Clause:    filter,
		Result:    &result,
	}
	err := repositories.NewRepository().GetBy(param)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetAll = Get all data of Setting
func (N *SettingService) GetAll(sortKey, sortBy string, skip, limit int, filter toolkit.M) ([]*models.SettingModel, int, error) {
	resultRows := make([]*models.SettingModel, 0)
	var resultTotal int
	param := repositories.GetAllParam{
		TableName:   models.NewSettingModel().CollName(),
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

// Insert = Insert New Setting Data
func (N *SettingService) Insert(data *models.SettingModel) error {
	if data.ID == "" {
		data.ID = bson.NewObjectId()
	}

	param := repositories.InsertParam{
		TableName: models.NewSettingModel().CollName(),
		Data:      data,
	}

	err := repositories.NewRepository().Insert(param)

	if err != nil {
		return err
	}

	return nil
}

// Update = Update existing data
func (N *SettingService) Update(data *models.SettingModel) error {
	param := repositories.SaveParam{
		TableName: models.NewSettingModel().CollName(),
		Data:      data,
	}

	err := repositories.NewRepository().Save(param)

	if err != nil {
		return err
	}

	return nil
}

// Delete = Delete Setting Data
func (N *SettingService) Delete(clause *dbflex.Filter) error {
	param := repositories.DeleteParam{
		TableName: models.NewSettingModel().CollName(),
		Clause:    clause,
	}

	err := repositories.NewRepository().Delete(param)

	if err != nil {
		return err
	}

	return nil
}

// Find = Find Setting Data
func (N *SettingService) Find(clause *dbflex.Filter) (bool, []*models.SettingModel, error) {
	res, err := N.Get(clause)
	if err != nil {
		return false, nil, err
	}

	if len(res) > 0 {
		return true, res, nil
	}

	return false, nil, nil
}
