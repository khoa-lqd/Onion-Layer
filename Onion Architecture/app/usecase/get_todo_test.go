package usecase

import (
	"OnionPractice/app/domain/model"
	mockRepo "OnionPractice/util/testhelper/mock/repository"
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestTodoUseCase_Get(t *testing.T) {
	now := time.Now()
	todo := model.NewTodo("name", false, now, now, now)

	type fields struct {
		setTodoRepository func(mock *mockRepo.MockTodoRepository)
	}

	type args struct {
		input GetTodoUseCaseInput
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				setTodoRepository: func(mock *mockRepo.MockTodoRepository) {
					mock.EXPECT().Get(gomock.Any(), gomock.Any())
				},
			},
			args: args{
				input: GetTodoUseCaseInput{
					ID: todo.ID(),
				},
			},
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			todoRepository := mockRepo.NewMockTodoRepository(ctrl)
			test.fields.setTodoRepository(todoRepository)

			u := GetTodoUseCase{
				todoRepository: todoRepository,
			}
			res, err := u.Get(context.Background(), test.args.input)
			if test.wantErr {
				assert.NotNil(t, err)
				assert.Nil(t, res)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, res)
			}
		})
	}
}
