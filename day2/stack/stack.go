package stack

import "fmt"
type Stack[T any] struct {
    items []T
}

// Push adds an item to the stack
func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

// Pop removes and returns the top item of the stack
// Returns an error if the stack is empty
func (s *Stack[T]) Pop() (T, error) {
    if len(s.items) == 0 {
        var zero T // Return zero value of T
        return zero, fmt.Errorf("stack is empty")
    }
    // Get the last element
    top := s.items[len(s.items)-1]
    // Remove the last element
    s.items = s.items[:len(s.items)-1]
    return top, nil
}

// Peek returns the top item without removing it
// Returns an error if the stack is empty
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

func (s *Stack[T]) Copy() Stack[T] {
    newItems := make([]T, len(s.items))
    copy(newItems, s.items) // Copy the underlying slice
    return Stack[T]{items: newItems}
}
