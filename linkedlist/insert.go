package linkedlist

func (l *Node[T]) Insert(index int, i T) {
	found := l.Index(index)

	node := &Node[T]{
		i,
		found.tail,
	}
	found.tail = node
}
