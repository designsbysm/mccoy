package linkedlist

func (l *Node[T]) Reverse() *Node[T] {
	list := Empty[T]()

	l.ForEach(func(node *Node[T]) {
		list = list.Cons(node.head)
	})

	return list
}
