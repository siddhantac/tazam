package store

import "tazam/task"

// Store defines the interface for any storage implementation
type Store interface {
	Create(t task.Task) (int, error)
	Read(id int) (task.Task, error)
	Update(t, value task.Task) error
	Delete(t task.Task) error
	List() ([]task.Task, error)
}
