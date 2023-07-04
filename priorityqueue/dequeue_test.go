package priorityqueue

import "testing"

func TestDequeue(t *testing.T) {
	queue := createQueue()

	if length := queue.Length(); length != 6 {
		t.Errorf("should have 6, got %d", length)
	}

	item := queue.Dequeue()

	if item.head != "item5" {
		t.Errorf("should be item5, got %v", item.head)
	}

	if length := queue.Length(); length != 5 {
		t.Errorf("should have 5, got %d", length)
	}
}
