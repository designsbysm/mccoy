package priorityqueue

func (q *Queue[T]) ForEach(fn func(node *Node[T])) {
	q.lock.Lock()
	defer q.lock.Unlock()

	for node := q.nodes; node != nil; node = node.tail {
		fn(node)
	}
}
