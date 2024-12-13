package main

import (
	"fmt"
	"time"
)

type status int

const (
	todo status = iota
	inProgress
	done
)

func (s status) String() string {
	return [...]string{"todo", "in progress", "done"}[s]
}

/*
A note on SQL statements:
Make sure you're using parameterized SQL statements to avoid
SQL injections. This format creates prepared statements at run time.
learn more: https://go.dev/doc/database/sql-injection
*/

// note for reflect: only exported fields of a struct are settable.
type task struct {
	ID      uint
	Name    string
	Project string
	Status  string
	Created time.Time
}

func (t task) String() string {
	return fmt.Sprintf("name: %s\n project: %s\n status: %s\n created: %s", t.Name, t.Project, t.Status, t.Created.Format("2006-01-02"))
}

// implement list.Item & list.DefaultItem
func (t task) FilterValue() string {
	return t.Name
}

func (t task) Title() string {
	return t.Name
}

func (t task) Description() string {
	return t.Project
}

// implement kancli.Status
func (s status) Next() int {
	if s == done {
		return int(todo)
	}
	return int(s + 1)
}

func (s status) Prev() int {
	if s == todo {
		return int(done)
	}
	return int(s - 1)
}

func (s status) Int() int {
	return int(s)
}