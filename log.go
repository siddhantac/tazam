package main

import (
	"fmt"
	"tazam/task"
)

func logOperation(task task.Task, log string) {
	fmt.Printf("task %d %s\n", task.ID, log)
}
