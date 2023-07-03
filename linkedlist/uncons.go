package linkedlist

func (l *Node[T]) Uncons() (T, *Node[T]) {
	// TODO: does the original l stay in memory
	return l.head, l.tail
}
