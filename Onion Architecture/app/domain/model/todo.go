package model

import "time"

const UndefinedID = 0

type Todo struct {
	id        int
	name      string
	complete  bool
	deadline  time.Time
	createdAt time.Time
	updatedAt time.Time
}

func (e *Todo) ID() int {
	return e.id
}

func (e *Todo) Name() string {
	return e.name
}

func (e *Todo) Complete() bool {
	return e.complete
}

func (e *Todo) Deadline() time.Time {
	return e.deadline
}

func (e *Todo) CreatedAt() time.Time {
	return e.createdAt
}

func (e *Todo) UpdatedAt() time.Time {
	return e.updatedAt
}

func (e *Todo) IsOverdue() bool {
	if e.complete {
		return false
	}

	return time.Now().After(e.deadline)
}

func (e *Todo) ExtendDealine(extension time.Duration) {
	e.deadline = e.deadline.Add(extension)
}

func NewTodo(
	name string,
	complete bool,
	deadline time.Time,
	createdAt time.Time,
	updatedAt time.Time,
) Todo {
	return Todo{
		id:        UndefinedID,
		name:      name,
		complete:  complete,
		deadline:  deadline,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func RestoreTodo(
	id int,
	name string,
	complete bool,
	deadline time.Time,
	createdAt time.Time,
	updatedAt time.Time,
) Todo {
	return Todo{
		id:        id,
		name:      name,
		complete:  complete,
		deadline:  deadline,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}
