package main

import "fmt"
import "log"

var currentId int

var todos Todos

func init() {
	RepoCreateTodo(Todo{Name: "Write presentation"})
	RepoCreateTodo(Todo{Name: "Host meetup"})
}

func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.Id == id {
			log.Println(t)
			return t
		}
	}
	//return empty Todo if not found
	return Todo{}
}

func RepoCreateTodo(t Todo) Todo {
	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	log.Println(t)
	return t
}

func RepoDestroyTodo(id int) error {
	for index, t := range todos {
		if t.Id == id {
			todos = append(todos[:index], todos[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
