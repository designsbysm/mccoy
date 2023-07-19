package priorityqueue

import "sync"

type Queue[T any] struct {
	length int
	lock   sync.RWMutex
	nodes  *Node[T]
}

type Node[T any] struct {
	priorty int
	head    T
	tail    *Node[T]
}
