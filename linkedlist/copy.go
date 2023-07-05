package linkedlist

func (l *Node[T]) Copy() *Node[T] {
	list := Empty[T]()

	l.ForEach(func(node *Node[T]) {
		list = list.Cons(node.head)
	})

	return list.Reverse()
}
