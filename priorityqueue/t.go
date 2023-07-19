package priorityqueue

type Queue[T any] struct {
	length int
	nodes  *Node[T]
}

type Node[T any] struct {
	priorty int
	head    T
	tail    *Node[T]
}
