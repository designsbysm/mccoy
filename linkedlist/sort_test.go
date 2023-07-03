package linkedlist

import (
	"strconv"
	"testing"

	"github.com/designsbysm/mccoy"
)

func TestSort(t *testing.T) {
	array := []mccoy.Item{31, 13, 12, 4, 18, 16, 7, 2, 3, 0, 10}
	list := FromArray(array)
	sorted := list.Sort()

	var got string
	for node := sorted; node != nil; node = node.tail {
		got += strconv.Itoa(node.head.(int))
	}

	should := "02347101213161831"
	if should != got {
		t.Errorf("should have %s, got %s", should, got)
	}
}
