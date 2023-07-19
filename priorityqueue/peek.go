package priorityqueue

func (q *Queue[T]) Peek() T {
	return q.nodes.head
}
