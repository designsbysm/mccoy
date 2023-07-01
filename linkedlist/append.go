package linkedlist

import "github.com/designsbysm/mccoy"

func (n *T) Append(i mccoy.Item) {
	node := New(i)
	last := n.Last()

	if last != nil {
		last.tail = node
	}
}
