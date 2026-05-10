package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/vhula/cli-tools/todo"
)

const todoFileName = ".todo.json"

func main() {
	todoList := &todo.TodoList{}
	if err := todoList.Get(todoFileName); err != nil {
		log.Fatalln(err)
	}
	switch {
	case len(os.Args) == 1:
		for _, item := range *todoList {
			fmt.Fprintln(os.Stdout, item.Task)
		}
	default:
		item := strings.Join(os.Args[1:], " ")
		todoList.Add(item)
		if err := todoList.Save(todoFileName); err != nil {
			log.Fatalln(err)
		}
	}
}
