package linkedlist

func (l *Node[T]) ToArray() []T {
	array := []T{}

	for node := l; node != nil; node = node.tail {
		array = append(array, node.head)
	}

	return array
}
