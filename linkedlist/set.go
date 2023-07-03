package linkedlist

func (l *Node[T]) Set(index int, i T) {
	found := l.Index(index)
	found.head = i
}
