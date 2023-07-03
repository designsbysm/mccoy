package linkedlist

func (l *T) Last() *T {
	// TODO: check empty list?, return error
	last := Empty()

	for node := l; node != nil; node = node.tail {
		if node.tail == nil {
			last = node
		}
	}

	return last
}
