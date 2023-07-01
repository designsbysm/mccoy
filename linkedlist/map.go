package linkedlist

import "github.com/designsbysm/mccoy"

func (n *T) Map(fn func(mccoy.Item) mccoy.Item) *T {
	list := Empty()

	for node := n; node != nil; node = node.tail {
		if node.head == nil {
			break
		}

		value := fn(node.head)

		if list.head == nil {
			list.head = value
		} else {
			list.Append(value)
		}
	}

	return list
}
