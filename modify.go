package main

import (
	"flag"
	"tazam/task"
)

func modifyTask(t task.Task, args []string) (task.Task, error) {
	if len(args) == 0 { // status update
		t.Status = t.Status.Next()
		logOperation(t, "Updated status to "+t.Status.String())
		return t, nil
	}

	modifyFS := flag.NewFlagSet("modify", flag.ExitOnError)

	name := modifyFS.String("name", t.Name, "Task name")
	project := modifyFS.String("project", t.Project, "project")
	priority := modifyFS.Int("priority", t.Priority, "priority")

	if err := modifyFS.Parse(args); err != nil {
		return task.Task{}, err
	}

	t.Name = *name
	t.Project = *project
	t.Priority = *priority

	logOperation(t, "updated")
	return t, nil
}
