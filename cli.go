package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
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
		taskTable(tasks)
		// for _, t := range tasks {
		// 	fmt.Printf("%v\n", t)
		// }
	case modifyCmd:
		fmt.Println("modify")
	case archiveCmd:
		fmt.Println("archive")
	default:
		return fmt.Errorf("unknown command")
	}

	return nil
}

func taskTable(tasks []task) {

	tbl := table.New().
		// BorderColumn(false).
		// Border(lipgloss.NormalBorder()).
		// BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("239"))).
		BorderHeader(true).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == -1:
				return lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
			case row%2 == 0:
				return lipgloss.NewStyle().Background(lipgloss.Color("232"))
			default:
				return lipgloss.NewStyle().Background(lipgloss.Color("234"))
			}
		}).
		Headers("ID", "Task", "Status", "Priority", "Project")

	for _, t := range tasks {
		tbl.Row(fmt.Sprintf("%d", t.ID), t.Name, t.Status, fmt.Sprintf("%d", t.Priority), t.Project)
	}
	fmt.Println(tbl)
}
