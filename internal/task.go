package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
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

func (t *TaskManager) LoadTasks() error {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	return json.Unmarshal(data, t)
}

func (t *TaskManager) SaveTasks() error {
	data, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("tasks.json", data, 0644)

}

func (t *TaskManager) Add(title string) {
	if err := t.LoadTasks(); err != nil {
		fmt.Println("load error: ", err)
		return
	}

	t.Tasks = append(t.Tasks, Task{
		ID:        t.NextId,
		Title:     title,
		Done:      false,
		CreatedAt: time.Now(),
	})
	t.NextId++

	if err := t.SaveTasks(); err != nil {
		fmt.Println("save error: ", err)
		return
	}
	fmt.Println("task added successfully")

}

func (t *TaskManager) List() {
	if err := t.LoadTasks(); err != nil {
		fmt.Println("error loading tasks: ", err)
		return
	}
	if len(t.Tasks) == 0 {
		fmt.Println("no tasks")
		return
	}

	for _, task := range t.Tasks {
		status := "[ ]"
		if task.Done {
			status = "[x]"
		}
		fmt.Printf("%s %d: %s \n", status, task.ID, task.Title)
	}
}

func (t *TaskManager) Done(id int) {
	if err := t.LoadTasks(); err != nil {
		fmt.Println("load error: ", err)
		return
	}
	for i := range t.Tasks {
		if t.Tasks[i].ID == id {
			t.Tasks[i].Done = true
			t.SaveTasks()
			fmt.Println("task done")
			return
		}
	}
	fmt.Println("task not found")
}

func (t *TaskManager) Delete(id int) {
	if err := t.LoadTasks(); err != nil {
		fmt.Println("error loading: ", err)
		return
	}
	for i := range t.Tasks {
		if t.Tasks[i].ID == id {
			t.Tasks = append(t.Tasks[:i], t.Tasks[i+1:]...)
			t.SaveTasks()
			fmt.Println("task deleted successfully")
			return
		}
	}
	fmt.Println("task not found")
}
