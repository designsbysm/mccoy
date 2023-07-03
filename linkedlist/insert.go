package linkedlist

import "github.com/designsbysm/mccoy"

func (l *T) Insert(index int, i mccoy.Item) {
	// TODO: change to At
	found := l.Index(index)

	node := &T{i, found.tail}
	found.tail = node
}
