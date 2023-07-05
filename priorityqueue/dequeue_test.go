package priorityqueue

import "testing"

func TestDequeue(t *testing.T) {
	queue := createQueue()

	if length := queue.Length(); length != 5 {
		t.Errorf("should have 5, got %d", length)
	}

	item := queue.Dequeue()

	if item.head != "item4" {
		t.Errorf("should be item4, got %v", item.head)
	}

	if length := queue.Length(); length != 4 {
		t.Errorf("should have 4, got %d", length)
	}
}
