package linkedlist

import "github.com/designsbysm/mccoy"

func (l *T) Prepend(i mccoy.Item) *T {
	return l.Cons(i)
}
