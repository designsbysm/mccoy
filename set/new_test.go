package set

import "testing"

func TestNew(t *testing.T) {
	set := New[string]()

	if set.items == nil {
		t.Errorf("should not be nil")
	}

	if length := set.Length(); length != 0 {
		t.Errorf("should be empty, got %d", length)
	}
}
