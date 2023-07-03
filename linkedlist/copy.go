package linkedlist

func (l *T) Copy() *T {
	list := New(l.head)

	for node := l.tail; node != nil; node = node.tail {
		list = list.Prepend(node.head)
	}

	return list.Reverse()
}
