package main

import "fmt"

func logOperation(task Task, log string) {
	fmt.Printf("task %d %s\n", task.ID, log)
}
