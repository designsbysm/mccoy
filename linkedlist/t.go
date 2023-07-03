package linkedlist

type Node[T any] struct {
	head T
	tail *Node[T]
}
