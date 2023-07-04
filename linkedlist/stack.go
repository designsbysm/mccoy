package linkedlist

// Aliases for LIFO (last in first out) stack

func (l *Node[T]) Push(i T) *Node[T] {
	return l.Cons(i)
}

func (l *Node[T]) Pop() (T, *Node[T]) {
	return l.Uncons()
}
