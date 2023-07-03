package linkedlist

import "github.com/designsbysm/mccoy"

func (l *T) Filter(fn func(mccoy.Item) bool) *T {
	list := Empty()

	for node := l; node != nil; node = node.tail {
		if node.head == nil {
			break
		} else if !fn(node.head) {
			continue
		}

		if list.head == nil {
			list.head = node.head
		} else {
			list.Append(node.head)
		}
	}

	return list
}
