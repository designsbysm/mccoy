package linkedlist

func (n *T) Size() int {
	counter := 0

	for node := n; node != nil; node = node.tail {
		counter += 1
	}

	return counter
}
