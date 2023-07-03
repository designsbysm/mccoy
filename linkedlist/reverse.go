package linkedlist

func (l *Node[T]) Reverse() *Node[T] {
	list := New(l.head)

	for node := l.tail; node != nil; node = node.tail {
		list = list.Cons(node.head)
	}

	return list
}
