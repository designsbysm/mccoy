package linkedlist

func FromArray[T any](array []T) *Node[T] {
	var list *Node[T]

	for _, i := range array {
		if list == nil {
			list = New(i)
		} else {
			list = list.Cons(i)
		}
	}

	return list.Reverse()
}
