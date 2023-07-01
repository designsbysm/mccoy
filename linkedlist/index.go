package linkedlist

func (n *T) Index(index int) *T {
	// TODO: check bounds, return error
	node := n

	for i := 0; i < index; i++ {
		tail := node.tail

		if tail != nil {
			node = tail
		}
	}

	return node
}
