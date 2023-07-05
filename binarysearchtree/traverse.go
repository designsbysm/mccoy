package binarysearchtree

func (t *Tree[T]) Traverse(fn func(*Node[T])) {
	t.lock.RLock()
	defer t.lock.RUnlock()

	traverseTree(t.root, fn)
}

func traverseTree[T any](node *Node[T], fn func(*Node[T])) {
	if node != nil {
		traverseTree(node.left, fn)
		fn(node)
		traverseTree(node.right, fn)
	}
}
