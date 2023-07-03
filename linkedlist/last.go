package linkedlist

func (l *Node[T]) Last() *Node[T] {
	var last *Node[T]

	for node := l; node != nil; node = node.tail {
		if node.tail == nil {
			last = node
		}
	}

	return last
}
