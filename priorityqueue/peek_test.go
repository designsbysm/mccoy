package priorityqueue

import "testing"

func TestPeek(t *testing.T) {
	queue := createQueue()
	item := queue.Peek()

	if item.head != "item4" {
		t.Errorf("should be item4, got %v", item.head)
	}

	if length := queue.Length(); length != 5 {
		t.Errorf("should have 5, got %d", length)
	}
}
