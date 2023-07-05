package priorityqueue

type Queue[T any] struct {
	nodes *Node[T]
}

type Node[T any] struct {
	priorty int
	head    T
	tail    *Node[T]
}
