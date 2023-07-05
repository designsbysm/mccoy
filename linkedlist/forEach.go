package linkedlist

func (l *Node[T]) ForEach(fn func(node *Node[T])) {
	for node := l; node != nil; node = node.tail {
		fn(node)
	}
}
