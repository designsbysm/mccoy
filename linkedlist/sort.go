package linkedlist

func (l *Node[T]) Sort(compare func(T, T) bool) *Node[T] {
	return split(
		l.Copy(),
		compare,
	)
}

func split[T any](left *Node[T], compare func(T, T) bool) *Node[T] {
	// don't need to split one or no items
	if left.tail == nil {
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
		split(left, compare),
		split(right, compare),
		compare,
	)
}

func merge[T any](left *Node[T], right *Node[T], compare func(T, T) bool) *Node[T] {
	list := Empty[T]()
	var item T

	for left != nil && right != nil {
		if compare(left.head, right.head) {
			item, left = left.Uncons()
			list = list.Cons(item)
		} else {
			item, right = right.Uncons()
			list = list.Cons(item)
		}
	}

	// prepend leftover items
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
