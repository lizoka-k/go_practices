package main

import (
	"fmt"
	"less3/list"
	"less3/queue"
	"less3/stack"
)

func main() {
	fmt.Println("=== Задание 1: Стек на дженериках ===")
	s := stack.NewStack[int](5)
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Printf("Вершина стека: %v\n", s.Peek())
	fmt.Printf("Извлечено: %v\n", s.Pop())
	fmt.Printf("Вершина стека: %v\n", s.Peek())
	fmt.Printf("Стек пуст? %v\n", s.IsEmpty())

	fmt.Println("\n=== Задание 2: Очередь на дженериках ===")
	q := queue.NewQueue[string](5)
	q.Push("первый")
	q.Push("второй")
	q.Push("третий")
	fmt.Printf("Извлечено из очереди: %v\n", q.Pop())
	fmt.Printf("Извлечено из очереди: %v\n", q.Pop())
	fmt.Printf("Очередь пуста? %v\n", q.IsEmpty())

	fmt.Println("\n=== Задание 3: Односвязный список на дженериках ===")
	l := list.NewSinglyLinkedList[int]()
	l.Add(100)
	l.Add(200)
	l.Add(300)
	fmt.Printf("Элемент по индексу 1: %v\n", l.Get(1))
	l.Remove(1)
	fmt.Printf("Все значения: %v\n", l.Values())
	fmt.Printf("Размер списка: %v\n", l.Size())
}
