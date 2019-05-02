package repositories

import (
	"log"
	"notes/app/core"
	"reflect"

	"git.eaciitapp.com/sebar/dbflex"
	"github.com/eaciit/toolkit"
)

// Copyright 2018 Noval Agung Prayogo. All rights reserved.

type Repository struct {
	BaseRepository
}

func NewRepository() *Repository {
	return new(Repository)
}

type GetAllParam struct {
	Filter      toolkit.M
	SortKey     string
	SortOrder   string
	Skip        int
	Take        int
	TableName   string
	ResultRows  interface{}
	ResultTotal *int
}

func (Repository) GetAll(param GetAllParam) error {
	sortKey := param.SortKey
	if sortKey == "" {
		sortKey = "CreatedAt"
	}

	sortOrder := -1
	if param.SortKey == "asc" {
		sortOrder = 1
	}

	pipeRows := make([]toolkit.M, 0)

	if len(param.Filter) > 0 {
		pipeRows = append(pipeRows, toolkit.M{"$match": param.Filter})
	}

	pipeRows = append(pipeRows, toolkit.M{"$sort": toolkit.M{sortKey: sortOrder}})

	if param.Skip > 0 {
		pipeRows = append(pipeRows, toolkit.M{"$skip": param.Skip})
	}
	if param.Take > 0 {
		pipeRows = append(pipeRows, toolkit.M{"$limit": param.Take})
	}

	// if core.IsDev() {
	log.Println("query", toolkit.JsonString(pipeRows))
	// }

	err := core.Database().Cursor(
		dbflex.From(param.TableName).Command("pipe"),
		toolkit.M{"pipe": pipeRows},
	).Fetchs(param.ResultRows, 0)
	if err != nil {
		return err
	}

	// ------------ count

	pipeTotal := make([]toolkit.M, 0)

	if len(param.Filter) > 0 {
		pipeTotal = append(pipeTotal, toolkit.M{"$match": param.Filter})
	}

	pipeTotal = append(pipeTotal, toolkit.M{
		"$group": toolkit.M{
			"_id": nil,
			"total": toolkit.M{
				"$sum": 1,
			},
		},
	})

	if core.IsDev() {
		log.Println("query", toolkit.JsonString(pipeRows))
	}

	resultTotal := make([]toolkit.M, 0)
	err = core.Database().Cursor(
		dbflex.From(param.TableName).Command("pipe"),
		toolkit.M{"pipe": pipeTotal},
	).Fetchs(&resultTotal, 0)
	if err != nil {
		return err
	}

	if len(resultTotal) > 0 && param.ResultTotal != nil {
		*param.ResultTotal = resultTotal[0].GetInt("total")
	}

	return nil
}

type GetByParam struct {
	TableName string
	Clause    *dbflex.Filter
	Result    interface{}
}

func (Repository) GetBy(param GetByParam) error {
	query := dbflex.From(param.TableName)
	if param.Clause != nil {
		query = query.Where(param.Clause)
	}

	cursor := core.Database().Cursor(
		query.Select(),
		nil)
	err := cursor.Fetchs(param.Result, 0)
	if err != nil {
		return err
	}

	return nil
}

type AggrParam struct {
	TableName string
	Pipe      []toolkit.M
	Result    interface{}
}

func (Repository) Aggr(param AggrParam) error {
	if core.IsDev() {
		log.Println("query", toolkit.JsonString(param.Pipe))
	}

	err := core.Database().Cursor(
		dbflex.From(param.TableName).Command("pipe"),
		toolkit.M{"pipe": param.Pipe},
	).Fetchs(param.Result, 0)
	if err != nil {
		return err
	}

	return nil
}

type InsertParam struct {
	TableName       string
	Data            interface{}
	ContinueOnError bool
}

func (Repository) Insert(param InsertParam) error {
	query, err := core.Database().Prepare(
		dbflex.From(param.TableName).Insert(),
	)
	if err != nil {
		return err
	}

	var lastError error

	switch reflect.TypeOf(param.Data).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(param.Data)

		for i := 0; i < s.Len(); i++ {
			_, err = query.Execute(toolkit.M{}.Set("data", s.Index(i).Interface()))
			lastError = err
			if !param.ContinueOnError && err != nil {
				return err
			}
		}
	default:
		_, err = query.Execute(toolkit.M{}.Set("data", param.Data))
		lastError = err
	}

	return lastError
}

type SaveParam struct {
	TableName       string
	Data            interface{}
	ContinueOnError bool
}

func (Repository) Save(param SaveParam) error {
	query, err := core.Database().Prepare(
		dbflex.From(param.TableName).Save(),
	)
	if err != nil {
		return err
	}

	var lastError error

	switch reflect.TypeOf(param.Data).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(param.Data)

		for i := 0; i < s.Len(); i++ {
			_, err = query.Execute(toolkit.M{}.Set("data", s.Index(i).Interface()))
			lastError = err
			if !param.ContinueOnError && err != nil {
				return err
			}
		}
	default:
		_, err = query.Execute(toolkit.M{}.Set("data", param.Data))
		lastError = err
	}

	return lastError
}

type DeleteParam struct {
	TableName string
	Clause    *dbflex.Filter
}

func (Repository) Delete(param DeleteParam) error {
	query := dbflex.From(param.TableName)
	if param.Clause != nil {
		query = query.Where(param.Clause)
	}
	query = query.Delete()

	_, err := core.Database().Execute(query, nil)
	return err
}
