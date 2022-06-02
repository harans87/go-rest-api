package main

import "fmt"

var currentId int

var todos Todos

func init() {
	// createTodo()
}

func createTodo(t Todo) Todo {
	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	return t
}

func findTodo(id int) Todo {
	for _, t1 := range todos {
		if t1.Id == id {
			return t1
		}
	}
	return Todo{}
}

func deleteTodo(id int) error {
	for i, t1 := range todos {
		if t1.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("could not find the todo")
}
