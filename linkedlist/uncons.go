package linkedlist

import "github.com/designsbysm/mccoy"

func (l *T) Uncons() (mccoy.Item, *T) {
	// TODO: does the original l stay in memory
	return l.head, l.tail
}
