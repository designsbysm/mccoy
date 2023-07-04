package linkedlist

func (l *Node[T]) Insert(index int, head T) {
	found := l.Index(index)

	node := &Node[T]{
		head: head,
		tail: found.tail,
	}

	found.tail = node
}
