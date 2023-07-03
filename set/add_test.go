package set

import (
	"testing"
)

func TestAdd(t *testing.T) {
	set := New()

	if size := set.Length(); size != 0 {
		t.Errorf("should be empty, got %d", size)
	}

	set.Add("item1")
	if size := set.Length(); size != 1 {
		t.Errorf("should have 1, got %d", size)
	}

	set.Add("item1")
	if size := set.Length(); size != 1 {
		t.Errorf("should have 1, got %d", size)
	}

	set.Add("item2")
	if size := set.Length(); size != 2 {
		t.Errorf("should have 2, got %d", size)
	}
}

func TestAddEmpty(t *testing.T) {
	set := T{}

	set.Add("item1")
	if size := set.Length(); size != 1 {
		t.Errorf("should have 1, got %d", size)
	}
}
