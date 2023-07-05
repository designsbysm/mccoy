package priorityqueue

import "testing"

func TestEnqueueEmpty(t *testing.T) {
	queue := New[string]()
	queue.Enqueue(5, "item2")

	if length := queue.Length(); length != 1 {
		t.Errorf("should have 1, got %d", length)
	}
}

func TestEnqueueMultiple(t *testing.T) {
	queue := New[string]()
	queue.Enqueue(5, "item1")
	queue.Enqueue(1, "item2")
	queue.Enqueue(9, "item3")
	queue.Enqueue(7, "item4")
	queue.Enqueue(6, "item5")
	queue.Enqueue(5, "item6")
	queue.Enqueue(5, "item7")

	if length := queue.Length(); length != 7 {
		t.Errorf("should have 7, got %d", length)
	}

	inOrder := true
	queue.ForEach(func(node *Node[string]) {
		if node.tail == nil {
			return
		}

		inOrder = node.priorty <= node.tail.priorty
	})

	if !inOrder {
		t.Errorf("should be in order, got %v", toArray(queue))
	}
}
