package main

import (
	"fmt"
	"strconv"
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
		if len(args) > 1 {
			id, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return err
			}

			task, err := db.getTask(uint(id))
			if err != nil {
				return err
			}
			fmt.Println(task)
			return nil
		}

		tasks, err := db.getTasks()
		if err != nil {
			return err
		}
		taskTable(tasks)

	case modifyCmd:
		if len(args) < 2 {
			return fmt.Errorf("too few arguments")
		}
		if len(args) == 2 {
			id, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return err
			}

			task, err := db.getTask(uint(id))
			if err != nil {
				return err
			}
			task = modifyTask(task, args[1:])
			if err := db.update(task); err != nil {
				return err
			}
			fmt.Printf("Task %d: %s\n", task.ID, task.Status)
			return nil
		}
		return fmt.Errorf("unimplemented")
	case archiveCmd:
		fmt.Println("archive")
	default:
		return fmt.Errorf("unknown command")
	}

	return nil
}

func modifyTask(task task, args []string) task {
	if len(args) == 1 {
		s := StatusFromString(task.Status)
		task.Status = s.Next().String()
	}
	return task
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
