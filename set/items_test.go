package set

import "testing"

func TestItems(t *testing.T) {
	set := New[string]()

	set.Add("item1")
	set.Add("item2")
	set.Add("item3")
	items := set.Items()

	if len(items) != 3 {
		t.Errorf("should have 3, got %d", len(items))
	}

	found := items[1]
	if found != "item2" {
		t.Errorf("should have item2, got %s", found)
	}
}
