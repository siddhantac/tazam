package main

import (
	"database/sql"
	"fmt"
	"log"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	path := "/Users/siddhanta.c/workspace/projects/tazam/tasks"
	if err := initTaskDir(path); err != nil {
		log.Fatal(err)
	}

	db, err := openDB(path)
	if err != nil {
		log.Fatal(err)
	}
	if !db.tableExists() {
		if err := db.createTable(); err != nil {
			log.Fatal(err)
		}
	}

	if err := db.insert("task 1", "project 1"); err != nil {
		log.Printf("err: failed to insert task: %v", err)
	}
	tasks, err := db.getTasks()
	if err != nil {
		log.Printf("err: failed to get tasks: %v", err)
	}
	for _, task := range tasks {
		fmt.Printf("task: %v\n\n", task)
	}
}

// openDB opens a SQLite database and stores that database in our special spot.
func openDB(path string) (*taskDB, error) {
	db, err := sql.Open("sqlite3", filepath.Join(path, "tasks.db"))
	if err != nil {
		return nil, err
	}
	t := taskDB{db, path}
	if !t.tableExists() {
		err := t.createTable()
		if err != nil {
			return nil, err
		}
	}
	return &t, nil
}
