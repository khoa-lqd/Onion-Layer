package database

import (
	"OnionPractice/app/domain/model"
	"OnionPractice/app/domain/repository"
	"OnionPractice/app/infrastructure/database/dbmodel"
	"context"

	"gorm.io/gorm"
)

type TodoRepositoryImpl struct {
	db *gorm.DB
}

func (r TodoRepositoryImpl) Create(ctx context.Context, todo model.Todo) (*model.Todo, error) {
	m := r.domainModelToDBModel(todo) // từ domain xuống DB
	if err := r.db.Create(m).Error; err != nil {
		return nil, err
	}

	return r.dbModelToDomainModelPointer(*m), nil // khi return thì return pointer từ db tới domain
}

func (r TodoRepositoryImpl) Get(ctx context.Context, id int) (*model.Todo, error) {
	var todo dbmodel.Todo
	if err := r.db.First(&todo, id).Error; err != nil {
		return nil, err
	}

	return r.dbModelToDomainModelPointer(todo), nil
}

func (r TodoRepositoryImpl) domainModelToDBModel(entity model.Todo) *dbmodel.Todo {
	return &dbmodel.Todo{
		ID:        entity.ID(),
		Name:      entity.Name(),
		Complete:  entity.Complete(),
		Deadline:  entity.Deadline(),
		CreatedAt: entity.CreatedAt(),
		UpdatedAt: entity.UpdatedAt(),
	}
}

func (r TodoRepositoryImpl) dbModelToDomainModel(todo dbmodel.Todo) model.Todo {
	return model.RestoreTodo(
		todo.ID,
		todo.Name,
		todo.Complete,
		todo.Deadline,
		todo.CreatedAt,
		todo.UpdatedAt,
	)
}

func (r TodoRepositoryImpl) dbModelToDomainModelPointer(todo dbmodel.Todo) *model.Todo {
	e := r.dbModelToDomainModel(todo)

	return &e
}

func NewTodoRepository(db *gorm.DB) repository.TodoRepository {
	return &TodoRepositoryImpl{db: db}
}
