package set

import "testing"

func TestAdd(t *testing.T) {
	set := New[string]()

	if length := set.Length(); length != 0 {
		t.Errorf("should be 0, got %d", length)
	}

	set.Add("item1")
	set.Add("item2")
	if size := set.Length(); size != 2 {
		t.Errorf("should have 2, got %d", size)
	}
}

func TestAddEmpty(t *testing.T) {
	set := New[string]()
	set.items = nil

	if length := set.Length(); length != 0 {
		t.Errorf("should be 0, got %d", length)
	}

	set.Add("item1")
	set.Add("item2")
	if size := set.Length(); size != 2 {
		t.Errorf("should have 2, got %d", size)
	}
}
