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
	queue := createQueue()

	inOrder := true
	for node := queue.nodes; node != nil; node = node.tail {
		if node.tail == nil {
			break
		}

		inOrder = node.priorty <= node.tail.priorty
	}

	if !inOrder {
		t.Errorf("should be in order, got %v", toArray(queue))
	}
}
