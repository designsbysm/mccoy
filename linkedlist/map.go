package linkedlist

func (l *Node[T]) Map(fn func(T) T) *Node[T] {
	list := Empty[T]()

	for node := l; node != nil; node = node.tail {
		value := fn(node.head)

		if list == nil {
			list = New(value)
		} else {
			list.Append(value)
		}
	}

	return list
}
