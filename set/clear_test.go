package set

import (
	"testing"
)

func TestClear(t *testing.T) {
	set := New()

	set.Add("item1")
	set.Add("item2")
	set.Add("item3")
	set.Clear()

	if size := set.Size(); size != 0 {
		t.Errorf("should be empty, got %d", size)
	}
}
