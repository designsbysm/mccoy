package priorityqueue

import "testing"

func TestNew(t *testing.T) {
	queue := New[string]()

	if queue.nodes != nil {
		t.Errorf("should be nil, got %v", queue.nodes)
	}
}
