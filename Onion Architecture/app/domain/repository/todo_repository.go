package repository

import (
	"OnionPractice/app/domain/model"
	"context"
)

type TodoRepository interface {
	Create(ctx context.Context, todo model.Todo) (*model.Todo, error)
	Get(ctx context.Context, id int) (*model.Todo, error)
}
