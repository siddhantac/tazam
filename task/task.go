package task

import (
	"fmt"
	"time"
)

/*
A note on SQL statements:
Make sure you're using parameterized SQL statements to avoid
SQL injections. This format creates prepared statements at run time.
learn more: https://go.dev/doc/database/sql-injection
*/

// note for reflect: only exported fields of a struct are settable.
type Task struct {
	ID       int
	Name     string
	Priority int
	Project  string
	Status   status
	Created  time.Time
}

func New(name string) Task {
	return Task{
		Name:    name,
		Created: time.Now(),
		Status:  todo,
	}
}

func (t Task) String() string {
	return fmt.Sprintf("%s\n project: %s\n status: %s\n created: %s", t.Name, t.Project, t.Status, t.Created.Format("2006-01-02"))
}

// implement list.Item & list.DefaultItem
func (t Task) FilterValue() string {
	return t.Name
}

func (t Task) Title() string {
	return t.Name
}

func (t Task) Description() string {
	return t.Project
}
