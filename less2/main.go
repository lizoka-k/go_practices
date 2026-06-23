package main

import "fmt"

func main() {
	fmt.Println("=== Задание 1: Стек ===")
	s := newStack(5)
	push(s, 10)
	push(s, 20)
	push(s, 30)
	fmt.Printf("Вершина стека: %v\n", peek(s))
	fmt.Printf("Извлечено: %v\n", pop(s))
	fmt.Printf("Вершина стека: %v\n", peek(s))

	fmt.Println("\n=== Задание 2: Очередь ===")
	q := newQueue(5)
	pushQueue(q, "первый")
	pushQueue(q, "второй")
	pushQueue(q, "третий")
	fmt.Printf("Извлечено из очереди: %v\n", popQueue(q))
	fmt.Printf("Извлечено из очереди: %v\n", popQueue(q))

	fmt.Println("\n=== Задание 3: Односвязный список ===")
	list := newSinglyLinkedList()
	add(list, 100)
	add(list, 200)
	add(list, 300)
	fmt.Printf("Элемент по индексу 1: %v\n", get(list, 1))
	remove(list, 1)
	fmt.Printf("Все значения: %v\n", values(list))

	fmt.Println("\n=== Задание 4: Конвертер римских цифр ===")
	fmt.Printf("XIV = %d\n", romanToArabic("XIV"))
	fmt.Printf("MCMXCIX = %d\n", romanToArabic("MCMXCIX"))
	fmt.Printf("XLII = %d\n", romanToArabic("XLII"))

	fmt.Println("\n=== Задание 5: Двумерный массив с уникальными числами ===")
	matrix := fillUniqueMatrix(3, 4)
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
}
