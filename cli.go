package main

import (
	"fmt"
	"strings"
)

type cmd string

const (
	addCmd     cmd = "add"
	listCmd    cmd = "list"
	modifyCmd  cmd = "modify"
	archiveCmd cmd = "archive"
)

// processCmds processes command like args and flags
// format: ./tazam cmd taskID flags
func processCmds(args []string, db *taskDB) error {
	// var userCmdStr string
	// flag.StringVar(&userCmdStr, "cmd", "", "Command to execute")
	// flag.Parse()

	userCmd := cmd(args[0])

	switch userCmd {
	case addCmd:
		t := newTask(strings.Join(args[1:], " "))
		id, err := db.insert(t)
		if err != nil {
			return err
		}
		fmt.Printf("created task %d\n", id)
	case listCmd:
		tasks, err := db.getTasks()
		if err != nil {
			return err
		}
		for _, t := range tasks {
			fmt.Printf("%v\n", t)
		}
	case modifyCmd:
		fmt.Println("modify")
	case archiveCmd:
		fmt.Println("archive")
	default:
		return fmt.Errorf("unknown command")
	}

	return nil
}
