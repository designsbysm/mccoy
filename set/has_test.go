package set

import "testing"

func TestHas(t *testing.T) {
	set := New[string]()

	set.Add("item1")
	set.Add("item2")
	set.Add("item3")

	if ok := set.Has("item1"); !ok {
		t.Errorf("should have item1")
	}

	if ok := set.Has("item4"); ok {
		t.Errorf("should not have item4")
	}
}
