package main

import (
	"flag"
	"strings"
	"tazam/store"
	"tazam/task"
)

const defaultPriority = 2

func addTask(args []string, db store.Store) (task.Task, error) {
	addFS := flag.NewFlagSet("add", flag.ExitOnError)
	name := addFS.String("name", "", "task.Task name")
	project := addFS.String("project", "", "project")
	priority := addFS.Int("priority", defaultPriority, "priority")
	addFS.Parse(args[1:])

	var t task.Task

	if *name != "" {
		t = task.New(*name)
	} else {
		t = task.New(strings.Join(args[1:], " "))
	}

	if *project != "" {
		t.Project = *project
	}

	t.Priority = *priority

	id, err := db.Create(t)
	if err != nil {
		return task.Task{}, err
	}
	t.ID = int(id)
	return t, nil
}
