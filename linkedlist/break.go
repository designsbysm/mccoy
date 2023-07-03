package linkedlist

func (l *Node[T]) Break(index int) (*Node[T], *Node[T]) {
	left := l.Copy()

	found := left.Index(index)
	right := found.tail
	found.tail = nil

	return left, right
}
