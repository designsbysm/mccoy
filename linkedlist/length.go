package linkedlist

func (l *T) Length() int {
	counter := 0

	for node := l; node != nil; node = node.tail {
		counter += 1
	}

	return counter
}
