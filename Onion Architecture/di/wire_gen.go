// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"OnionPractice/app/infrastructure/database"
	"OnionPractice/app/usecase"
	"OnionPractice/db"
	"github.com/google/wire"
)

// Injectors from define.go:

// use case
func CreateTodoUseCase() (usecase.CreateTodoUseCase, error) {
	gormDB := db.GetDB()
	todoRepository := database.NewTodoRepository(gormDB)
	createTodoUseCase := usecase.NewCreateTodoUseCase(todoRepository)
	return createTodoUseCase, nil
}

func GetTodoUseCase() (usecase.GetTodoUseCase, error) {
	gormDB := db.GetDB()
	todoRepository := database.NewTodoRepository(gormDB)
	getTodoUseCase := usecase.NewGetTodoUseCase(todoRepository)
	return getTodoUseCase, nil
}

// define.go:

var providerSet = wire.NewSet(db.GetDB, database.NewTodoRepository)