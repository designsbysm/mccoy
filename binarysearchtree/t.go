package binarysearchtree

import "sync"

type Tree[T any] struct {
	root *Node[T]
	lock sync.RWMutex
}

type Node[T any] struct {
	key   int
	value T
	left  *Node[T]
	right *Node[T]
}
