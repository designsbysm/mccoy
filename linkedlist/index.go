package linkedlist

func (l *Node[T]) Index(index int) *Node[T] {
	// TODO: check bounds, return error
	node := l

	for i := 0; i < index; i++ {
		tail := node.tail

		if tail != nil {
			node = tail
		}
	}

	return node
}
