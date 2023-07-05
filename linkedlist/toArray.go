package linkedlist

func (l *Node[T]) ToArray() []T {
	array := []T{}

	l.ForEach(func(node *Node[T]) {
		array = append(array, node.head)
	})

	return array
}
