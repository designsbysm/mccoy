package linkedlist

func New[T any](i T) *Node[T] {
	return &Node[T]{
		i,
		nil,
	}
}
