package priorityqueue

import (
	"fmt"
)

func createQueue() *Queue[string] {
	queue := New[string]()
	queue.Enqueue(5, "item1")
	queue.Enqueue(1, "item2")
	queue.Enqueue(9, "item3")
	queue.Enqueue(-3, "item4")
	queue.Enqueue(7, "item5")

	return queue
}

func toArray(q *Queue[string]) []string {
	array := []string{}

	for node := q.nodes; node != nil; node = node.tail {
		str := fmt.Sprintf("%d:%s", node.priorty, node.head)
		array = append(array, str)
	}

	return array
}
