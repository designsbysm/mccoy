package linkedlist

import "github.com/designsbysm/mccoy"

func (n *T) Set(index int, i mccoy.Item) {
	found := n.Index(index)
	found.head = i
}
