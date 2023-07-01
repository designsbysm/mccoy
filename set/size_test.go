package set

import (
	"testing"
)

func TestSzie(t *testing.T) {
	set := New()

	set.Add("item1")
	set.Add(2)
	set.Add(3.14)
	items := set.Items()

	if len(items) != 3 {
		t.Errorf("should have 3, got %d", len(items))
	}
}
