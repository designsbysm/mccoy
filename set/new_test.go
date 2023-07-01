package set

import (
	"testing"
)

func TestNew(t *testing.T) {
	set := New()

	if set.items == nil {
		t.Errorf("should not be nil")
	}

	if size := set.Size(); size != 0 {
		t.Errorf("should be empty, got %d", size)
	}
}
