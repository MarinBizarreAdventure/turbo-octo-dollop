package main

import (
	"fmt"
	"os"
	"strconv"

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
		if len(os.Args) < 3 {
			fmt.Println("usage: add <title>")
			return
		}
		tm.Add(os.Args[2])
	case "list":
		tm.List()
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("usage: done <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ascii to int error: ", err)
			return
		}
		tm.Done(id)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("usage: delete <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ascii to int error: ", err)
			return
		}
		tm.Delete(id)
	default:
		fmt.Println("unknown command:", command)
	}
}
