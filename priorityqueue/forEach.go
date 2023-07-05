package priorityqueue

func (q *Queue[T]) ForEach(fn func(node *Node[T])) {
	for node := q.nodes; node != nil; node = node.tail {
		fn(node)
	}
}
