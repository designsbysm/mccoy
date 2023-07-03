package linkedlist

func (l *T) Delete(index int) {
	found := l.Index(index - 1)
	next := found.tail

	if next != nil {
		found.tail = next.tail
	}
}
