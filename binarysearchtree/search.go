package binarysearchtree

func (t *Tree[T]) Search(key int) bool {
	t.lock.RLock()
	defer t.lock.RUnlock()

	return searchTree(t.root, key)
}

func searchTree[T any](node *Node[T], key int) bool {
	if node == nil {
		return false
	} else if key < node.key {
		return searchTree(node.left, key)
	} else if key > node.key {
		return searchTree(node.right, key)
	}

	return true
}
