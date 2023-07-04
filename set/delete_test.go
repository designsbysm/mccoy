package set

import "testing"

func TestDelete(t *testing.T) {
	set := New[string]()

	set.Add("item1")
	set.Add("item2")
	set.Add("item3")
	set.Delete("item2")

	if length := set.Length(); length != 2 {
		t.Errorf("should have 2, got %d", length)
	}

	if ok := set.Has("item2"); ok {
		t.Errorf("should not have item2")
	}
}
