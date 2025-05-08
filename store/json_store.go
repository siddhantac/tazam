package store

import (
	"encoding/json"
	"fmt"
	"os"
	"tazam/task"
)

// JSONStore implements Store interface using a JSON file
type JSONStore struct {
	filename string
	data     map[int]task.Task
}

// NewJSONStore creates a new JSONStore instance
func NewJSONStore(filename string) (*JSONStore, error) {
	store := &JSONStore{
		filename: filename,
		data:     make(map[int]task.Task),
	}

	// Create file if it doesn't exist
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		if err := store.save(); err != nil {
			return nil, fmt.Errorf("failed to create store file: %w", err)
		}
	} else {
		// Load existing data
		if err := store.load(); err != nil {
			return nil, fmt.Errorf("failed to load store file: %w", err)
		}
	}

	return store, nil
}

// load reads the JSON file into memory
func (s *JSONStore) load() error {
	file, err := os.ReadFile(s.filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	if len(file) == 0 {
		s.data = make(map[int]task.Task)
		return nil
	}

	if err := json.Unmarshal(file, &s.data); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return nil
}

// save writes the in-memory data to the JSON file
func (s *JSONStore) save() error {
	data, err := json.MarshalIndent(s.data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	if err := os.WriteFile(s.filename, data, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// Create adds a new key-value pair to the store
func (s *JSONStore) Create(t task.Task) (int, error) {
	t.ID = len(s.data) + 1
	s.data[t.ID] = t
	return t.ID, s.save()
}

// Read retrieves a value from the store by key
func (s *JSONStore) Read(id int) (task.Task, error) {
	var t task.Task
	data, exists := s.data[id]
	if !exists {
		return t, fmt.Errorf("key not found: %v", t.ID)
	}

	return data, nil
}

// Update modifies an existing key-value pair in the store
func (s *JSONStore) Update(t task.Task) error {
	if _, exists := s.data[t.ID]; !exists {
		return fmt.Errorf("key not found: %v", t.ID)
	}

	s.data[t.ID] = t
	return s.save()
}

// Delete removes a key-value pair from the store
func (s *JSONStore) Delete(t task.Task) error {
	if _, exists := s.data[t.ID]; !exists {
		return fmt.Errorf("key not found: %v", t.ID)
	}

	delete(s.data, t.ID)
	return s.save()
}

// List returns all key-value pairs in the store
func (s *JSONStore) List() ([]task.Task, error) {
	// Create a copy of the data to prevent external modification
	result := make([]task.Task, len(s.data))
	for k, v := range s.data {
		result[k-1] = v
	}

	return result, nil
}
