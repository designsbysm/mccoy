package linkedlist

// Aliases for FIFO (first in first out) queue

func (l *Node[T]) Enqueue(i T) {
	l.Append(i)
}

func (l *Node[T]) Dequeue() (T, *Node[T]) {
	return l.Uncons()
}
