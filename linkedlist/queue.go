package linkedlist

// Aliases for FIFO (first in first out) queue

func (l *Node[T]) Enqueue(head T) {
	l.Append(head)
}

func (l *Node[T]) Dequeue() (T, *Node[T]) {
	return l.Uncons()
}
