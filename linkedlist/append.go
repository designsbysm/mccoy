package linkedlist

import "github.com/designsbysm/mccoy"

func (l *T) Append(i mccoy.Item) {
	node := New(i)
	last := l.Last()

	if last != nil {
		last.tail = node
	}
}
