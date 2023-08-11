package integration__test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"testing"
)

func TestTodoScenarion(t *testing.T) {
	ctx := context.Background()

	// CreateTodo
	createTodoRes, err := createTodo(ctx)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(createTodoRes)

	// GetTodo
	getTodoRes, err := getTodo(ctx, createTodoRes.ID)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(getTodoRes)

}

type CreateTodoResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Complete  bool   `json:"complete"`
	Deadline  string `json:"deadline"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func createTodo(ctx context.Context) (*CreateTodoResponse, error) {
	baseURL := "http://localhost" + testServerPort + "/todos"

	msg := map[string]interface{}{
		"name": "todo1",
	}

	b, err := json.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to create todo: %w", err)
	}

	buf := new(bytes.Buffer)
	buf.Write(b)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, baseURL, buf)
	if err != nil {
		return nil, fmt.Errorf("failed to create todo: %w", err)
	}

	client := new(http.Client)

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to create todo: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to create todo: %w", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to create todo: %w", err)
	}

	var parsedRes CreateTodoResponse
	if err := json.Unmarshal(body, &parsedRes); err != nil {
		return nil, fmt.Errorf("failed to create todo: %w", err)
	}

	return &parsedRes, nil
}

type GetTodoResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Complete  bool   `json:"complete"`
	Deadline  string `json:"deadline"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func getTodo(ctx context.Context, id int) (*GetTodoResponse, error) {
	baseURL := "http://localhost" + testServerPort + "/todos/" + strconv.Itoa(id)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create todo: %w", err)
	}

	client := new(http.Client)

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to create todo: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to create todo: %w", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to create todo: %w", err)
	}

	var parsedRes GetTodoResponse
	if err := json.Unmarshal(body, &parsedRes); err != nil {
		return nil, fmt.Errorf("failed to create todo: %w", err)
	}

	return &parsedRes, nil
}
