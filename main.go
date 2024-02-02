package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"sync"
)

var baseURL = "https://jsonplaceholder.typicode.com/todos/"

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func fetchTodoByID(todoID int) (Todo, error) {
	resp, err := http.Get(baseURL + fmt.Sprint(todoID))
	if err != nil {
		return Todo{}, err
	}
	defer resp.Body.Close()

	var todo Todo
	if err := json.NewDecoder(resp.Body).Decode(&todo); err != nil {
		return Todo{}, err
	}

	return todo, nil
}

func printTodoDetails(todo Todo) {
	fmt.Printf("userId\t%d\n", todo.UserID)
	fmt.Printf("id\t%d\n", todo.ID)
	fmt.Printf("title\t%q\n", todo.Title)
	fmt.Printf("completed\t%v\n", todo.Completed)
	fmt.Println()
}


func setBaseURL(url string) {
	baseURL = url
}

func main() {
	var limit int
	flag.IntVar(&limit, "limit", 20, "Limit of TODOs to fetch")
	flag.Parse()

	var wg sync.WaitGroup
	for i := 2; i <= limit*2; i += 2 {
		wg.Add(1)
		go func(todoID int) {
			defer wg.Done()
			todo, err := fetchTodoByID(todoID)
			if err != nil {
				fmt.Printf("Error fetching TODO at index %d: %v\n", todoID/2, err)
				return
			}
			printTodoDetails(todo)
		}(i)
	}

	wg.Wait()
}
