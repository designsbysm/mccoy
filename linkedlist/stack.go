package linkedlist

// Aliases for LIFO (last in first out) stack

func (l *Node[T]) Push(head T) *Node[T] {
	return l.Cons(head)
}

func (l *Node[T]) Pop() (T, *Node[T]) {
	return l.Uncons()
}
