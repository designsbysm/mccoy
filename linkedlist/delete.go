package linkedlist

func (l *Node[T]) Delete(index int) {
	// TODO: check index bounds

	found := l.Index(index - 1)
	next := found.tail

	if next != nil {
		found.tail = next.tail
	}
}
