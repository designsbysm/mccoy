package linkedlist

func (l *Node[T]) Last() *Node[T] {
	var last *Node[T]

	l.ForEach(func(node *Node[T]) {
		if node.tail == nil {
			last = node
		}
	})

	return last
}
