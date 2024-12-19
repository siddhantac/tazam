package main

import (
	"flag"
)

func modifyTask(task Task, args []string) (Task, error) {
	if len(args) == 0 { // status update
		s := StatusFromString(task.Status)
		task.Status = s.Next().String()
		logOperation(task, "Updated status to "+task.Status)
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

	logOperation(task, "updated")
	return task, nil
}
