package priorityqueue

func (q *Queue[T]) Length() int {
	length := 0

	q.ForEach(func(node *Node[T]) {
		length++
	})

	return length
}
