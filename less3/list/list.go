package list

type Item[T any] struct {
	v    T
	next *Item[T]
}

type SinglyLinkedList[T any] struct {
	first *Item[T]
	last  *Item[T]
	size  int
}

func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

func (l *SinglyLinkedList[T]) Add(v T) {
	newItem := &Item[T]{v: v}
	if l.first == nil {
		l.first = newItem
		l.last = newItem
	} else {
		l.last.next = newItem
		l.last = newItem
	}
	l.size++
}

func (l *SinglyLinkedList[T]) Get(idx int) T {
	var zero T
	if idx < 0 || idx >= l.size {
		return zero
	}
	current := l.first
	for i := 0; i < idx; i++ {
		current = current.next
	}
	return current.v
}

func (l *SinglyLinkedList[T]) Remove(idx int) {
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

func (l *SinglyLinkedList[T]) Values() []T {
	result := make([]T, l.size)
	current := l.first
	for i := 0; i < l.size; i++ {
		result[i] = current.v
		current = current.next
	}
	return result
}

func (l *SinglyLinkedList[T]) Size() int {
	return l.size
}
