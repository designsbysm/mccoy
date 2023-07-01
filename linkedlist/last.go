package linkedlist

func (n *T) Last() *T {
	// TODO: check empty list?, return error
	last := Empty()

	for node := n; node != nil; node = node.tail {
		if node.tail == nil {
			last = node
		}
	}

	return last
}
