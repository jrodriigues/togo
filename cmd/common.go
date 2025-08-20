package cmd

import (
	"fmt"
	"github.com/prime-run/togo/model"
	"os"
)

func loadTodoListOrExit() *model.TodoList {
	todoList, err := model.LoadTodoListWithSource(TodoFileName, sourceFlag)
	if err != nil {
		fmt.Println("Error loading todos:", err)
		os.Exit(1)
	}
	return todoList
}

func saveTodoListOrExit(todoList *model.TodoList) {
	if err := todoList.SaveWithSource(TodoFileName, sourceFlag); err != nil {
		fmt.Println("Error saving todos:", err)
		os.Exit(1)
	}
}

func checkEmptyTodoList(todoList *model.TodoList, emptyMessage string) bool {
	if len(todoList.Todos) == 0 {
		fmt.Println(emptyMessage)
		return true
	}
	return false
}

func handleErrorAndExit(err error, message string) {
	if err != nil {
		fmt.Println(message, err)
		os.Exit(1)
	}
}
