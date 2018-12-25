package store

import (
	"fmt"
)

// Blackboard implements store.Interface and allows for
// reading and writing arbitrary data.
type Blackboard struct {
	data map[string]interface{}
}

// NewBlackboard returns a new blackboard.
func NewBlackboard() *Blackboard {
	s := &Blackboard{}
	s.data = make(map[string]interface{})
	return s
}

// Read returns the data associated with key if it exists,
// otherwise an error.
func (s *Blackboard) Read(key string) (interface{}, error) {
	value, ok := s.data[key]
	if !ok {
		return nil, fmt.Errorf("no data found for key %q", key)
	}
	return value, nil
}

// Write associates a key with some data.
func (s *Blackboard) Write(key string, data interface{}) {
	s.data[key] = data
}
