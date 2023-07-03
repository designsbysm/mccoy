package linkedlist

import "github.com/designsbysm/mccoy"

func (l *T) Cons(i mccoy.Item) *T {
	if l.head == nil {
		l.head = i
		return l
	} else {
		return &T{
			head: i,
			tail: l,
		}
	}
}
