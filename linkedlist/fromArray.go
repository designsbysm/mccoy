package linkedlist

import (
	"github.com/designsbysm/mccoy"
)

func FromArray(array []mccoy.Item) *T {
	list := Empty()

	for _, i := range array {
		if list.head != nil {
			list = list.Prepend(i)
		} else {
			list.head = i
		}
	}

	return list.Reverse()
}
