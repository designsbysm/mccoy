package linkedlist

func (l *Node[T]) Delete(index int) {
	found := l.Index(index - 1)
	next := found.tail

	if next != nil {
		found.tail = next.tail
	}
}
