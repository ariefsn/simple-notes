package services

import (
	"notes/app/models"
	"notes/app/repositories"

	"gopkg.in/mgo.v2/bson"

	"github.com/eaciit/toolkit"

	"git.eaciitapp.com/sebar/dbflex"
)

// CategoryService = Service for Category
type CategoryService struct {
}

// NewCategoryService = Create new service of Category
func NewCategoryService() *CategoryService {
	return new(CategoryService)
}

// Get = Get data of Category with dbflex filter
func (N *CategoryService) Get(Clause *dbflex.Filter) ([]*models.CategoryModel, error) {
	result := make([]*models.CategoryModel, 0)

	var filter *dbflex.Filter
	if Clause != nil {
		filter = Clause
	}

	param := repositories.GetByParam{
		TableName: models.NewCategoryModel().CollName(),
		Clause:    filter,
		Result:    &result,
	}
	err := repositories.NewRepository().GetBy(param)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetAll = Get all data of Category
func (N *CategoryService) GetAll(sortKey, sortBy string, skip, limit int, filter toolkit.M) ([]*models.CategoryModel, int, error) {
	resultRows := make([]*models.CategoryModel, 0)
	var resultTotal int
	param := repositories.GetAllParam{
		TableName:   models.NewCategoryModel().CollName(),
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

// Insert = Insert New Category Data
func (N *CategoryService) Insert(data *models.CategoryModel) error {
	if data.ID == "" {
		data.ID = bson.NewObjectId()
	}

	param := repositories.InsertParam{
		TableName: models.NewCategoryModel().CollName(),
		Data:      data,
	}

	err := repositories.NewRepository().Insert(param)

	if err != nil {
		return err
	}

	return nil
}

// Update = Update existing data
func (N *CategoryService) Update(data *models.CategoryModel) error {
	param := repositories.SaveParam{
		TableName: models.NewCategoryModel().CollName(),
		Data:      data,
	}

	err := repositories.NewRepository().Save(param)

	if err != nil {
		return err
	}

	return nil
}

// Delete = Delete Category Data
func (N *CategoryService) Delete(clause *dbflex.Filter) error {
	param := repositories.DeleteParam{
		TableName: models.NewCategoryModel().CollName(),
		Clause:    clause,
	}

	err := repositories.NewRepository().Delete(param)

	if err != nil {
		return err
	}

	return nil
}

// Find = Find Category Data
func (N *CategoryService) Find(clause *dbflex.Filter) (bool, []*models.CategoryModel, error) {
	res, err := N.Get(clause)
	if err != nil {
		return false, nil, err
	}

	if len(res) > 0 {
		return true, res, nil
	}

	return false, nil, nil
}
