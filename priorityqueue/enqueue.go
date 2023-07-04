package priorityqueue

func (q *Queue[T]) Enqueue(priorty int, head T) {
	if q.nodes == nil {
		q.nodes = &Node[T]{
			priorty: priorty,
			head:    head,
			tail:    nil,
		}
		q.length++

		return
	} else if priorty < q.nodes.priorty {
		q.nodes = &Node[T]{
			priorty: priorty,
			head:    head,
			tail:    q.nodes,
		}
		q.length++

		return
	}

	for node := q.nodes; node != nil; node = node.tail {
		if priorty <= node.priorty || node.tail == nil {
			node.tail = &Node[T]{
				priorty: priorty,
				head:    head,
				tail:    nil,
			}

			break
		} else if node.tail != nil && priorty < node.tail.priorty {
			node.tail = &Node[T]{
				priorty: priorty,
				head:    head,
				tail:    node.tail,
			}

			break
		}
	}

	q.length++
}
