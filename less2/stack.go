package main

import "fmt"

type stack struct {
	s    []any
	head int
}

func newStack(size int) *stack {
	return &stack{
		s:    make([]any, size),
		head: -1,
	}
}

func push(s *stack, v any) {
	if s.head >= len(s.s)-1 {
		fmt.Println("Стек переполнен")
		return
	}
	s.head++
	s.s[s.head] = v
}

func pop(s *stack) any {
	if s.head < 0 {
		fmt.Println("Стек пуст")
		return nil
	}
	v := s.s[s.head]
	s.s[s.head] = nil
	s.head--
	return v
}

func peek(s *stack) any {
	if s.head < 0 {
		return nil
	}
	return s.s[s.head]
}
