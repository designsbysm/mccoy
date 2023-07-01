package linkedlist

func (n *T) Reverse() *T {
	list := New(n.head)

	for node := n.tail; node != nil; node = node.tail {
		list = list.Prepend(node.head)
	}

	return list
}
