package main

import (
	"fmt"
	"strconv"
	"tazam/store"
	"tazam/task"

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
func processCmds(args []string, db2 store.Store) error {
	// var userCmdStr string
	// flag.StringVar(&userCmdStr, "cmd", "", "Command to execute")
	// flag.Parse()

	var userCmd cmd
	if len(args) == 0 {
		userCmd = listCmd
	} else {
		userCmd = cmd(args[0])
	}

	switch userCmd {
	case addCmd:
		t, err := addTask(args, db2)
		if err != nil {
			return err
		}
		logOperation(t, "added")
	case listCmd:
		if len(args) > 1 {
			id, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return err
			}

			t, err := db2.Read(int(id))
			// task, err := db.getTask(uint(id))
			if err != nil {
				return err
			}
			fmt.Println(t)
			return nil
		}

		tasks, err := db2.List()
		// tasks, err := db.getTasks()
		if err != nil {
			return err
		}
		taskTable(tasks)

	case modifyCmd:
		if len(args) < 2 {
			return fmt.Errorf("too few arguments")
		}
		id, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid id: %w", err)
		}

		task, err := db2.Read(int(id))
		// task, err := db.getTask(uint(id))
		if err != nil {
			return fmt.Errorf("read: %w", err)
		}
		task, err = modifyTask(task, args[2:])
		if err != nil {
			return fmt.Errorf("modify: %w", err)
		}
		if err := db2.Update(task); err != nil {
			return fmt.Errorf("update: %w", err)
		}
		return nil
	case archiveCmd:
		fmt.Println("archive")
	default:
		return fmt.Errorf("unknown command")
	}

	return nil
}

func taskTable(tasks []task.Task) {
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
		tbl.Row(fmt.Sprintf("%d", t.ID), t.Name, t.Status.String(), fmt.Sprintf("%d", t.Priority), t.Project)
	}
	fmt.Println(tbl)
}
