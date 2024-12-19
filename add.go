package main

import (
	"flag"
	"strings"
)

const defaultPriority = 2

func addTask(args []string, db *taskDB) (Task, error) {
	addFS := flag.NewFlagSet("add", flag.ExitOnError)
	name := addFS.String("name", "", "Task name")
	project := addFS.String("project", "", "project")
	priority := addFS.Int("priority", defaultPriority, "priority")
	addFS.Parse(args[1:])

	var t Task

	if *name != "" {
		t = newTask(*name)
	} else {
		t = newTask(strings.Join(args[1:], " "))
	}

	if *project != "" {
		t.Project = *project
	}

	t.Priority = *priority

	id, err := db.insert(t)
	if err != nil {
		return Task{}, err
	}
	t.ID = uint(id)
	return t, nil
}
