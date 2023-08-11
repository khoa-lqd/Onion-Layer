package usecase

import (
	"OnionPractice/app/domain/model"
	"OnionPractice/app/domain/repository"
	"context"
	"fmt"
	"time"
)

type CreateTodoUseCase struct {
	todoRepository repository.TodoRepository
}

func NewCreateTodoUseCase(todoRepository repository.TodoRepository) CreateTodoUseCase {
	return CreateTodoUseCase{todoRepository: todoRepository}
}

type CreateTodoUseCaseInput struct {
	Name     string
	Deadline time.Time
}

type CreateTodoUseCaseOutput struct {
	Todo *model.Todo
}

func (u CreateTodoUseCase) Create(ctx context.Context, input CreateTodoUseCaseInput) (*CreateTodoUseCaseOutput, error) {
	now := time.Now()
	todo := model.NewTodo(input.Name, false, input.Deadline, now, now)

	createTodo, err := u.todoRepository.Create(ctx, todo)
	if err != nil {
		return nil, fmt.Errorf("failed to todoRepository.Create: %w", err)
	}

	return &CreateTodoUseCaseOutput{Todo: createTodo}, nil
}
