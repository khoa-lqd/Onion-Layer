package controller

import (
	"OnionPractice/app/usecase"
	"OnionPractice/di"
	"OnionPractice/helpers"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateTodoRequest struct {
	Name string `json:"name" validate:"required"`
}

func Create(c *gin.Context) {
	req := CreateTodoRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, helpers.RenderError(http.StatusBadRequest))

		return
	}

	u, err := di.CreateTodoUseCase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.RenderError(http.StatusInternalServerError))

		return
	}

	input := usecase.CreateTodoUseCaseInput{
		Name:     req.Name,
		Deadline: time.Now().Add(24 * time.Hour),
	}

	result, err := u.Create(c, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.RenderError(http.StatusInternalServerError))

		return
	}

	c.JSON(201, gin.H{
		"id":         result.Todo.ID(),
		"name":       result.Todo.Name(),
		"complete":   result.Todo.Complete(),
		"deadline":   result.Todo.Deadline().String(),
		"created_at": result.Todo.CreatedAt().String(),
		"updated_at": result.Todo.UpdatedAt().String(),
	})
}

func Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.RenderError(http.StatusBadRequest))

		return
	}

	u, err := di.GetTodoUseCase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.RenderError(http.StatusInternalServerError))
		return
	}

	input := usecase.GetTodoUseCaseInput{
		ID: id,
	}

	todo, err := u.Get(c, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.RenderError(http.StatusInternalServerError))
	}

	c.JSON(200, gin.H{
		"id":         todo.ID(),
		"name":       todo.Name(),
		"complete":   todo.Complete(),
		"deadline":   todo.Deadline().String(),
		"created_at": todo.CreatedAt().String(),
		"updated_at": todo.UpdatedAt().String(),
	})

}
