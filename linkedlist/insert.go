package linkedlist

import "github.com/designsbysm/mccoy"

func (n *T) Insert(index int, i mccoy.Item) {
	// TODO: change to At
	found := n.Index(index)

	node := &T{i, found.tail}
	found.tail = node
}
