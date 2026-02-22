package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

type TaskManager struct {
	Tasks  []Task `json:"tasks"`
	NextId int    `json:"next_id"`
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		NextId: 0,
		Tasks:  []Task{},
	}
}

func (t *TaskManager) LoadTasks() (*TaskManager, error) {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			initialData := &TaskManager{}
			jsonData, _ := json.Marshal(initialData)
			err = os.WriteFile("tasks.json", jsonData, 0644)
			if err != nil {
				return nil, err
			}
			return initialData, nil
		}
		return nil, err
	}
	var tasks TaskManager
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}
	t.NextId = tasks.NextId

	return &tasks, nil
}

func (t *TaskManager) SaveTasks(tasks *TaskManager) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("tasks.json", data, 0644)

}

func (t *TaskManager) Add() {
	tasks, err := t.LoadTasks()

	if err != nil {
		fmt.Println("load task error: ", err)
		return
	}

	if len(os.Args) < 3 {
		fmt.Println("there is no task to add")
		return
	}
	title := os.Args[2]
	task := &Task{
		ID:        t.NextId,
		Title:     title,
		Done:      false,
		CreatedAt: time.Now(),
	}
	t.NextId++

	tasks.Tasks = append(tasks.Tasks, *task)
	tasks.NextId = t.NextId
	err = t.SaveTasks(tasks)
	if err != nil {
		fmt.Println("add error: ", err)
		return
	}
	fmt.Println("task added successfully")

}

func (t *TaskManager) List() {
	tasks, err := t.LoadTasks()
	if err != nil {
		fmt.Println("error loading tasks: ", err)
		return
	}
	fmt.Println(tasks)
}

func (t *TaskManager) Done() {
	tasks, err := t.LoadTasks()
	if err != nil {
		fmt.Println("load err: ", err)
		return
	}
	sid := os.Args[2]
	if sid == "" {
		fmt.Println("no id was given as an argument")
		return
	}
	id, err := strconv.Atoi(sid)
	if err != nil {
		fmt.Println("error converting string to int: ", err)
		return
	}
	for i := range tasks.Tasks {
		if tasks.Tasks[i].ID == id {
			tasks.Tasks[i].Done = true
			t.SaveTasks(tasks)
			return
		}
	}
}

func (t *TaskManager) Delete() {
	tasks, err := t.LoadTasks()
	if err != nil {
		fmt.Println("error loading: ", err)
		return
	}
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("error converting id: ", err)
		return
	}
	for i := range tasks.Tasks {
		if tasks.Tasks[i].ID == id {
			tasks.Tasks = append(tasks.Tasks[:i], tasks.Tasks[i+1:]...)
			t.SaveTasks(tasks)
			return
		}
	}
}
