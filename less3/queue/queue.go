package queue

import "fmt"

type Queue[T any] struct {
	s       []T
	low     int
	high    int
	size    int
	maxSize int
}

func NewQueue[T any](size int) *Queue[T] {
	return &Queue[T]{
		s:       make([]T, size),
		size:    0,
		maxSize: size,
		low:     0,
		high:    -1,
	}
}

func (q *Queue[T]) Push(v T) {
	if q.size >= q.maxSize {
		fmt.Println("Очередь переполнена")
		return
	}
	q.high = (q.high + 1) % q.maxSize
	q.s[q.high] = v
	q.size++
}

func (q *Queue[T]) Pop() T {
	var zero T
	if q.size == 0 {
		fmt.Println("Очередь пуста")
		return zero
	}
	v := q.s[q.low]
	q.s[q.low] = zero
	q.low = (q.low + 1) % q.maxSize
	q.size--
	return v
}

func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}
