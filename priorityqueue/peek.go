package priorityqueue

func (q *Queue[T]) Peek() *Node[T] {
	return q.nodes
}
