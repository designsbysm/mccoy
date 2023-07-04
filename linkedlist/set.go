package linkedlist

func (l *Node[T]) Set(index int, head T) {
	found := l.Index(index)
	found.head = head
}
