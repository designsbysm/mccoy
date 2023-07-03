package linkedlist

func (l *Node[T]) Filter(fn func(T) bool) *Node[T] {
	var list *Node[T]

	for node := l; node != nil; node = node.tail {
		if !fn(node.head) {
			continue
		}

		if list == nil {
			list = New(node.head)
		} else {
			list.Append(node.head)
		}
	}

	return list
}
