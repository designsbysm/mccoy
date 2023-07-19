package priorityqueue

func (q *Queue[T]) Enqueue(priorty int, head T) {
	q.length++

	if q.nodes == nil || priorty < q.nodes.priorty {
		q.nodes = &Node[T]{
			priorty: priorty,
			head:    head,
			tail:    q.nodes,
		}

		return
	}

	for node := q.nodes; node != nil; node = node.tail {
		if node.tail == nil {
			node.tail = &Node[T]{
				priorty: priorty,
				head:    head,
				tail:    nil,
			}

			break
		} else if priorty < node.tail.priorty && node.tail != nil {
			node.tail = &Node[T]{
				priorty: priorty,
				head:    head,
				tail:    node.tail,
			}

			break
		}
	}
}
