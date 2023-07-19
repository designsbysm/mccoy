package priorityqueue

func (q *Queue[T]) Dequeue() *Node[T] {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.length--

	node := q.nodes
	q.nodes = node.tail
	node.tail = nil

	return node
}
