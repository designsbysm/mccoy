package linkedlist

import (
	"github.com/designsbysm/mccoy"
)

func (l *T) ToArray() []mccoy.Item {
	array := []mccoy.Item{}

	for node := l; node != nil; node = node.tail {
		array = append(array, node.head)
	}

	return array
}
