package priorityqueue

import "testing"

func TestLength(t *testing.T) {
	queue := createQueue()

	if length := queue.Length(); length != 5 {
		t.Errorf("should have 5, got %d", length)
	}
}
