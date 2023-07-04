package linkedlist

func (l *Node[T]) Cons(head T) *Node[T] {
	return &Node[T]{
		head: head,
		tail: l,
	}
}
