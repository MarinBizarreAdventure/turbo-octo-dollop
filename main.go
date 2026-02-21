package main

import (
	"fmt"
	"os"

	"github.com/MarinBizarreAdventure/task-cli/internal"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: task [add|list|done|delete] [args]")
		return
	}
	command := os.Args[1]
	tm := internal.NewTaskManager()

	switch command {
	case "add":
		tm.Add()
	case "list":
		tm.List()
	case "done":
		tm.Done()
	case "delete":
		tm.Delete()
	default:
		fmt.Println("unkown command:", command)
	}
}
