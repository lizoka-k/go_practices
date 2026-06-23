package main

type item struct {
	v    any
	next *item
}

type singlyLinkedList struct {
	first *item
	last  *item
	size  int
}

func newSinglyLinkedList() *singlyLinkedList {
	return &singlyLinkedList{}
}

func add(l *singlyLinkedList, v any) {
	newItem := &item{v: v}
	if l.first == nil {
		l.first = newItem
		l.last = newItem
	} else {
		l.last.next = newItem
		l.last = newItem
	}
	l.size++
}

func get(l *singlyLinkedList, idx int) any {
	if idx < 0 || idx >= l.size {
		return nil
	}
	current := l.first
	for i := 0; i < idx; i++ {
		current = current.next
	}
	return current.v
}

func remove(l *singlyLinkedList, idx int) {
	if idx < 0 || idx >= l.size {
		return
	}
	if idx == 0 {
		l.first = l.first.next
		if l.first == nil {
			l.last = nil
		}
	} else {
		current := l.first
		for i := 0; i < idx-1; i++ {
			current = current.next
		}
		current.next = current.next.next
		if current.next == nil {
			l.last = current
		}
	}
	l.size--
}

func values(l *singlyLinkedList) []any {
	result := make([]any, l.size)
	current := l.first
	for i := 0; i < l.size; i++ {
		result[i] = current.v
		current = current.next
	}
	return result
}
