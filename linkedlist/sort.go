package linkedlist

import (
	"github.com/designsbysm/mccoy"
)

func (l *T) Sort() *T {
	return split(l.Copy())
}

func split(left *T) *T {
	// don't need to split one or no items
	if left.Length() < 2 {
		return left
	}

	mid := left
	end := left

	// find the middle item
	for end.tail != nil {
		for i := 2; i != 0 && end.tail != nil; i-- {
			end = end.tail
		}

		if end.tail == nil {
			break
		}

		mid = mid.tail
	}

	// split the list at the middle
	right := mid.tail
	mid.tail = nil

	return merge(
		split(left),
		split(right),
	)
}

func merge(left *T, right *T) *T {
	list := Empty()
	var item mccoy.Item

	for left != nil && right != nil {
		if left.head.(int) < right.head.(int) {
			item, left = left.Uncons()
			list = list.Cons(item)
		} else {
			item, right = right.Uncons()
			list = list.Cons(item)
		}
	}

	// merge leftover items
	for left != nil {
		item, left = left.Uncons()
		list = list.Cons(item)
	}

	for right != nil {
		item, right = right.Uncons()
		list = list.Cons(item)
	}

	return list.Reverse()
}
