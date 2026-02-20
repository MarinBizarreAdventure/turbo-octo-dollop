package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: task [add|list|done|delete] [args]")
		return
	}

	command := os.Args[2]

	switch command {
	case "add":
		//TODO
	case "list":
		//TODO
	case "done":
		//TODO
	case "delete":
		//TODO
	default:
		fmt.Println("unkown command:", command)
	}
}
