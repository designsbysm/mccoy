package linkedlist

import "github.com/designsbysm/mccoy"

func (l *T) Set(index int, i mccoy.Item) {
	found := l.Index(index)
	found.head = i
}
