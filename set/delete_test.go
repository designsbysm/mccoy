package set

import (
	"testing"
)

func TestDelete(t *testing.T) {
	set := New()

	set.Add("item1")
	set.Add("item2")
	set.Add("item3")
	set.Delete("item2")

	if size := set.Size(); size != 2 {
		t.Errorf("should delete, expected 2 and got %d", size)
	}
}
