package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type TodoList []item

func (todoList *TodoList) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*todoList = append(*todoList, t)
}

func (todoList *TodoList) Complete(index int) error {
	if index <= 0 || index > len(*todoList) {
		return fmt.Errorf("Item %d does not exist", index)
	}
	task := &(*todoList)[index-1]
	task.Done = true
	task.CompletedAt = time.Now()
	return nil
}

func (todoList *TodoList) Delete(index int) error {
	if index <= 0 || index > len(*todoList) {
		return fmt.Errorf("Item %d does not exist", index)
	}
	list := *todoList
	*todoList = append(list[:index-1], list[index:]...)
	return nil
}

func (todoList *TodoList) Save(filename string) error {
	js, err := json.Marshal(todoList)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, js, 0644)
}

func (todoList *TodoList) Get(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return nil
	}
	return json.Unmarshal(file, todoList)
}
