package binarysearchtree

func (t *Tree[T]) Min() *Node[T] {
	t.lock.RLock()
	defer t.lock.RUnlock()

	node := t.root
	if node == nil {
		return nil
	}

	for {
		if node.left == nil {
			return node
		}

		node = node.left
	}
}
