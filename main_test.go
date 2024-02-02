package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchTodoByID(t *testing.T) {
	// Mocking HTTP server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := `{"userId": 1, "id": 1, "title": "delectus aut autem", "completed": false}`
		w.Write([]byte(resp))
	}))
	defer mockServer.Close()

	setBaseURL(mockServer.URL)

	todo, err := fetchTodoByID(1)
	if err != nil {
		t.Errorf("Error fetching TODO: %v", err)
	}

	expectedTodo := Todo{
		UserID:    1,
		ID:        1,
		Title:     "delectus aut autem",
		Completed: false,
	}

	if todo != expectedTodo {
		t.Errorf("Expected TODO %+v, got %+v", expectedTodo, todo)
	}
}

func TestPrintTodoDetails(t *testing.T) {
	var buf bytes.Buffer

	todo := Todo{
		UserID:    1,
		ID:        1,
		Title:     "delectus aut autem",
		Completed: false,
	}

	printTodoDetails(todo)

	expectedOutput := "userId\t1\n" +
		"id\t1\n" +
		"title\t\"delectus aut autem\"\n" +
		"completed\tfalse\n\n"

	if buf.String() != expectedOutput {
		t.Errorf("Print output incorrect. Expected:\n%s\nGot:\n%s", expectedOutput, buf.String())
	}
}

func TestMainFunction(t *testing.T) {
	var buf bytes.Buffer

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := `{"userId": 1, "id": 1, "title": "delectus aut autem", "completed": false}`
		w.Write([]byte(resp))
	}))
	defer mockServer.Close()

	setBaseURL(mockServer.URL)

	main()

	expectedOutput := "userId\t1\n" +
		"id\t1\n" +
		"title\t\"delectus aut autem\"\n" +
		"completed\tfalse\n\n"

	if buf.String() != expectedOutput {
		t.Errorf("Main output incorrect. Expected:\n%s\nGot:\n%s", expectedOutput, buf.String())
	}
}
