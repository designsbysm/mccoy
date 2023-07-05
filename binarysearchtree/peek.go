package binarysearchtree

func (t *Tree[T]) Peek() *Node[T] {
	return t.Min()
}
