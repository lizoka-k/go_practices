package main

import "fmt"

type queue struct {
	s       []any
	low     int
	high    int
	size    int
	maxSize int
}

func newQueue(size int) *queue {
	return &queue{
		s:       make([]any, size),
		size:    0,
		maxSize: size,
		low:     0,
		high:    -1,
	}
}

func pushQueue(q *queue, v any) {
	if q.size >= q.maxSize {
		fmt.Println("Очередь переполнена")
		return
	}
	q.high = (q.high + 1) % q.maxSize
	q.s[q.high] = v
	q.size++
}

func popQueue(q *queue) any {
	if q.size == 0 {
		fmt.Println("Очередь пуста")
		return nil
	}
	v := q.s[q.low]
	q.s[q.low] = nil
	q.low = (q.low + 1) % q.maxSize
	q.size--
	return v
}
