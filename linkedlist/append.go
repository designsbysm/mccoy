package linkedlist

func (l *Node[T]) Append(head T) {
	node := New(head)
	last := l.Last()

	if last != nil {
		last.tail = node
	}
}
