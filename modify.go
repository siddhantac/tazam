package main

import (
	"flag"
	"fmt"
)

func modifyTask(task Task, args []string) (Task, error) {
	if len(args) == 0 { // status update
		s := StatusFromString(task.Status)
		task.Status = s.Next().String()
		fmt.Printf("Updated task %d: %s\n", task.ID, task.Status)
		return task, nil
	}

	modifyFS := flag.NewFlagSet("modify", flag.ExitOnError)

	name := modifyFS.String("name", task.Name, "Task name")
	project := modifyFS.String("project", task.Project, "project")
	priority := modifyFS.Int("priority", task.Priority, "priority")

	if err := modifyFS.Parse(args); err != nil {
		return Task{}, err
	}

	task.Name = *name
	task.Project = *project
	task.Priority = *priority

	fmt.Printf("Updated task %v\n", task.ID)
	return task, nil
}
