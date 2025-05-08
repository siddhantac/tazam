package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"tazam/store"
	"tazam/tui"

	tea "github.com/charmbracelet/bubbletea"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// path := "/Users/siddhanta.c/workspace/projects/tazam/data"
	// if err := initTaskDir(path); err != nil {
	// 	log.Fatal(err)
	// }
	//
	// db, err := openDB(path)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.db.Close()

	// if !db.tableExists() {
	// 	if err := db.createTable(); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	jsondb, err := store.NewJSONStore("data/tasks.json")
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) > 1 && os.Args[1] == "tui" {
		tasks, err := jsondb.List()
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
		p := tea.NewProgram(tui.New(tasks, jsondb))
		if _, err := p.Run(); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
	}

	if err := processCmds(os.Args[1:], jsondb); err != nil {
		log.Fatal(err)
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
