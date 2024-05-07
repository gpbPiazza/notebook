package stacks

import "errors"

var (
	errStackOverFlow = errors.New("stack over flow")
)

type Stack[T any] struct {
	capacity int
	store    []T
}

func NewStack[T any](capacity int) *Stack[T] {
	return &Stack[T]{
		capacity: capacity,
		store:    make([]T, 0),
	}
}

func (s *Stack[T]) Push(data T) error {
	if len(s.store)+1 > s.capacity {
		return errStackOverFlow
	}
	s.store = append(s.store, data)
	return nil
}

func (s *Stack[T]) Pop() T {
	if len(s.store) == 0 {
		return s.zeroValue()
	}

	data := s.store[s.lastIndex()]

	s.store = s.store[:s.lastIndex()]

	return data
}

func (s *Stack[T]) Read() T {
	if len(s.store) == 0 {
		return s.zeroValue()
	}

	return s.store[s.lastIndex()]
}

func (s *Stack[T]) lastIndex() int {
	return len(s.store) - 1
}

func (s *Stack[T]) zeroValue() T {
	var zero T
	return zero
}
