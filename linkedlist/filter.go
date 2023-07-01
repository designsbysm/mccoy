package linkedlist

import "github.com/designsbysm/mccoy"

func (n *T) Filter(fn func(mccoy.Item) bool) *T {
	list := Empty()

	for node := n; node != nil; node = node.tail {
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
