package main

import (
	"database/sql"
	"fmt"
	"os"
)

type taskDB struct {
	db      *sql.DB
	dataDir string
}

func initTaskDir(path string) error {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return os.Mkdir(path, 0o770)
		}
		return err
	}
	return nil
}

func (t *taskDB) tableExists() bool {
	if _, err := t.db.Query("SELECT * FROM tasks"); err == nil {
		return true
	}
	return false
}

func (t *taskDB) createTable() error {
	_, err := t.db.Exec(`CREATE TABLE "tasks" ( "id" INTEGER, "name" TEXT NOT NULL, "project" TEXT, "status" TEXT, "priority" INTEGER, "created" DATETIME, PRIMARY KEY("id" AUTOINCREMENT))`)
	return err
}

func (t *taskDB) insert(task Task) (int64, error) {
	// We don't care about the returned values, so we're using Exec. If we
	// wanted to reuse these statements, it would be more efficient to use
	// prepared statements. Learn more:
	// https://go.dev/doc/database/prepared-statements
	result, err := t.db.Exec(
		"INSERT INTO tasks(name, project, status, priority, created) VALUES(?, ?, ?, ?, ?)",
		task.Name,
		task.Project,
		task.Status,
		task.Priority,
		task.Created)
	id, _ := result.LastInsertId()
	return id, err
}

func (t *taskDB) delete(id uint) error {
	_, err := t.db.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}

// Update the task in the db. Provide new values for the fields you want to
// change, keep them empty if unchanged.
func (t *taskDB) update(task Task) error {
	// Get the existing state of the task we want to update.
	// orig, err := t.getTask(task.ID)
	// if err != nil {
	// 	return err
	// }
	//orig.merge(task)
	_, err := t.db.Exec(
		"UPDATE tasks SET name = ?, project = ?, status = ?, priority = ? WHERE id = ?",
		task.Name,
		task.Project,
		task.Status,
		task.Priority,
		task.ID)
	return err
}

func (t *taskDB) getTasks() ([]Task, error) {
	var tasks []Task
	rows, err := t.db.Query("SELECT * FROM tasks")
	if err != nil {
		return tasks, fmt.Errorf("unable to get values: %w", err)
	}
	for rows.Next() {
		var task Task
		err = rows.Scan(
			&task.ID,
			&task.Name,
			&task.Project,
			&task.Status,
			&task.Priority,
			&task.Created,
		)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, task)
	}
	return tasks, err
}

func (t *taskDB) getTasksByStatus(status string) ([]Task, error) {
	var tasks []Task
	rows, err := t.db.Query("SELECT * FROM tasks WHERE status = ?", status)
	if err != nil {
		return tasks, fmt.Errorf("unable to get values: %w", err)
	}
	for rows.Next() {
		var task Task
		err = rows.Scan(
			&task.ID,
			&task.Name,
			&task.Project,
			&task.Status,
			&task.Priority,
			&task.Created,
		)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, task)
	}
	return tasks, err
}

func (t *taskDB) getTask(id uint) (Task, error) {
	var task Task
	err := t.db.QueryRow("SELECT * FROM tasks WHERE id = ?", id).
		Scan(
			&task.ID,
			&task.Name,
			&task.Project,
			&task.Status,
			&task.Priority,
			&task.Created,
		)
	return task, err
}
