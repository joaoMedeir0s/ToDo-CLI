package main

import (
	"fmt"
	"os"
)

func use() {
	fmt.Println(`Use:
	tasks add <title>		add a task
	tasks list 				list the tasks
	tasks done <id> 		mark as concluded
	tasks rm <id> 			remove a task`)
}

// Verifica se o len args é valido (tasks + comando)
func main() {
	if len(os.Args) < 2 {
		use()
		os.Exit(1)
	}

	// Args é um vetor de arguentos, com o que o Os vai
	// receber do terminal
	command := os.Args[1]
	args := os.Args[2:]

	var err error
	switch command {
	case "add":
		err = cmdAdd(args)
	case "list":
		err = cmdList()
	case "done":
		err = cmdDone(args)
	case "remove":
		err = cmdRemove(args)
	default:
		fmt.Fprintf(os.Stderr, "Command unknown: %s\n\n", command)
		use()
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		os.Exit(1)
	}

}
