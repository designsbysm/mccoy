package priorityqueue

func (q *Queue[T]) Dequeue() *Node[T] {
	node := q.nodes
	q.nodes = node.tail
	q.length--
	node.tail = nil

	return node
}