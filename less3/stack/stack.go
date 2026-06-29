package stack

import "fmt"

type Stack[T any] struct {
	s    []T
	head int
}

func NewStack[T any](size int) *Stack[T] {
	return &Stack[T]{
		s:    make([]T, size),
		head: -1,
	}
}

func (s *Stack[T]) Push(v T) {
	if s.head >= len(s.s)-1 {
		fmt.Println("Стек переполнен")
		return
	}
	s.head++
	s.s[s.head] = v
}

func (s *Stack[T]) Pop() T {
	var zero T
	if s.head < 0 {
		fmt.Println("Стек пуст")
		return zero
	}
	v := s.s[s.head]
	s.s[s.head] = zero
	s.head--
	return v
}

func (s *Stack[T]) Peek() T {
	var zero T
	if s.head < 0 {
		return zero
	}
	return s.s[s.head]
}

func (s *Stack[T]) IsEmpty() bool {
	return s.head < 0
}