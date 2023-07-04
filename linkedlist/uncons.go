package linkedlist

func (l *Node[T]) Uncons() (T, *Node[T]) {
	return l.head, l.tail
}
