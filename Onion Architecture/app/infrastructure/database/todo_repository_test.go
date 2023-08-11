package database

import (
	"OnionPractice/app/domain/model"
	"OnionPractice/app/infrastructure/database/dbmodel"
	"OnionPractice/db"
	"OnionPractice/util/testhelper"
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func newTodoRepositoryImpl() *TodoRepositoryImpl {
	return &TodoRepositoryImpl{db: db.DB}
}

func getAllTodos() []dbmodel.Todo {
	var todos []dbmodel.Todo
	if err := db.DB.Find(&todos).Error; err != nil {
		panic(err)
	}

	return todos
}

func insertTodoTestData(todo model.Todo) int {
	dbModel := TodoRepositoryImpl{}.domainModelToDBModel(todo)
	if err := db.DB.Create(&dbModel).Error; err != nil {
		panic(err)
	}

	return dbModel.ID
}

func TestTodoRepositoryImpl_Create(t *testing.T) {
	now := time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)
	todo := model.NewTodo("todo", false, now, now, now)

	type args struct {
		todo model.Todo
	}

	tests := []struct {
		name   string
		args   args
		expect *model.Todo
	}{
		{
			name: "success",
			args: args{
				todo: todo,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo := newTodoRepositoryImpl()
			_, err := repo.Create(context.Background(), test.args.todo)
			assert.Nil(t, err)
			stored := getAllTodos()

			if diff := cmp.Diff(repo.dbModelToDomainModel(stored[0]), test.args.todo, testhelper.ToDoCmpOptions...); diff != "" {
				t.Errorf("Compare value is mismatch (-v1 +v2):%s\n", diff)
			}
		})
	}
}

func TestTodoRepositoryImp_Get(t *testing.T) {
	now := time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)
	todo := model.NewTodo("todo", false, now, now, now)
	id := insertTodoTestData(todo)

	type args struct {
		id int
	}

	tests := []struct {
		name   string
		args   args
		expect *model.Todo
	}{
		{
			name: "success",
			args: args{
				id: id,
			},
			expect: &todo,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo := newTodoRepositoryImpl()
			result, err := repo.Get(context.Background(), test.args.id)
			assert.Nil(t, err)
			if diff := cmp.Diff(result, test.expect, testhelper.ToDoCmpOptions...); diff != "" {
				t.Errorf("Compare value is mismatch (-v1 +v2):%s\n", diff)
			}
		})
	}
}
