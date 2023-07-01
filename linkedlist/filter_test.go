package linkedlist

import (
	"testing"

	"github.com/designsbysm/mccoy"
)

func TestFilter(t *testing.T) {
	list := createList()
	filtered := list.Filter(func(i mccoy.Item) bool {
		return i.(int) <= 5
	})

	if size := filtered.Size(); size != 5 {
		t.Errorf("should have 5, got %d", size)
	}
}

func TestFilterEmpty(t *testing.T) {
	list := Empty()
	filtered := list.Filter(func(i mccoy.Item) bool {
		return i.(int) <= 5
	})

	if size := filtered.Size(); size != 1 {
		t.Errorf("should have 1, got %d", size)
	}
}
