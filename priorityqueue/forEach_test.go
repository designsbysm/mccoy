package priorityqueue

import (
	"strconv"
	"testing"
)

func TestForEach(t *testing.T) {
	queue := createQueue()

	var got string
	queue.ForEach(func(node *Node[string]) {
		if node.tail == nil {
			return
		}

		got += strconv.Itoa(node.priorty)
	})

	if got != "-3157" {
		t.Errorf("should have -3157, got %s", got)
	}
}
