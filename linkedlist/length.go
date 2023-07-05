package linkedlist

func (l *Node[T]) Length() int {
	length := 0

	l.ForEach(func(node *Node[T]) {
		length += 1
	})

	return length
}
