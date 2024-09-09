package stack

import (
	"fmt"
)

// Stack struct using a slice
type Stack[T any] struct {
	items []T
}

// Push adds an element onto the stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes the top element from the stack and returns it
// Returns an error if the stack is empty
func (s *Stack[T]) Pop() (T, error) {
	if len(s.items) == 0 {
		var zero T // Return zero value of type T if stack is empty
		return zero, fmt.Errorf("stack is empty")
	}
	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return item, nil
}

// Peek returns the top element without removing it
func (s *Stack[T]) Peek() (T, error) {
	if len(s.items) == 0 {
		var zero T
		return zero, fmt.Errorf("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

// IsEmpty checks if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of elements in the stack
func (s *Stack[T]) Size() int {
	return len(s.items)
}
