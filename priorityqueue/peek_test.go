package priorityqueue

import "testing"

func TestPeek(t *testing.T) {
	queue := createQueue()
	item := queue.Peek()

	if item.head != "item5" {
		t.Errorf("should be item5, got %v", item.head)
	}

	if length := queue.Length(); length != 6 {
		t.Errorf("should have 6, got %d", length)
	}
}
