package usecase

import (
	"OnionPractice/app/domain/model"
	"OnionPractice/app/domain/repository"
	"context"
	"fmt"
)

type GetTodoUseCase struct {
	todoRepository repository.TodoRepository
}

func NewGetTodoUseCase(todoRepository repository.TodoRepository) GetTodoUseCase {
	return GetTodoUseCase{todoRepository: todoRepository}
}

type GetTodoUseCaseInput struct {
	ID int
}

func (u GetTodoUseCase) Get(ctx context.Context, input GetTodoUseCaseInput) (*model.Todo, error) {
	todo, err := u.todoRepository.Get(ctx, input.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to todoRepository.Get: %w", err)
	}

	return todo, nil
}
