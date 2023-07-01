package linkedlist

import (
	"testing"

	"github.com/designsbysm/mccoy"
)

func TestMap(t *testing.T) {
	list := createList()
	mapped := list.Map(func(i mccoy.Item) mccoy.Item {
		return i.(int) + 1
	})

	before := list.Index(4)
	if before.head != 5 {
		t.Errorf("should have 5, got %d", before.head)
	}

	after := mapped.Index(4)
	if after.head != 6 {
		t.Errorf("should have 6, got %d", after.head)
	}
}

func TestMapEmpty(t *testing.T) {
	list := Empty()
	mapped := list.Map(func(i mccoy.Item) mccoy.Item {
		return i.(int) + 1
	})

	if mapped.head != nil {
		t.Errorf("should be <nil>, got %v", list.head)
	}
}
