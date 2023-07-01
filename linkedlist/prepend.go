package linkedlist

import "github.com/designsbysm/mccoy"

func (n *T) Prepend(i mccoy.Item) *T {
	return &T{
		head: i,
		tail: n,
	}
}
