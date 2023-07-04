package linkedlist

func New[T any](head T) *Node[T] {
	return &Node[T]{
		head: head,
		tail: nil,
	}
}
