package priorityqueue

import "testing"

func TestLength(t *testing.T) {
	queue := createQueue()

	if length := queue.Length(); length != 6 {
		t.Errorf("should have 6, got %d", length)
	}
}
