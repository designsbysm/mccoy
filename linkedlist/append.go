package linkedlist

func (l *Node[T]) Append(i T) {
	node := New(i)
	last := l.Last()

	if last != nil {
		last.tail = node
	}
}
