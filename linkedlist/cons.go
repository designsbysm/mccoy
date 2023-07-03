package linkedlist

func (l *Node[T]) Cons(i T) *Node[T] {
	return &Node[T]{
		head: i,
		tail: l,
	}
}
