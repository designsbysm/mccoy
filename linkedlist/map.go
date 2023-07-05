package linkedlist

func (l *Node[T]) Map(fn func(T) T) *Node[T] {
	list := Empty[T]()

	l.ForEach(func(node *Node[T]) {
		value := fn(node.head)

		if list == nil {
			list = New(value)
		} else {
			list.Append(value)
		}
	})

	return list
}
