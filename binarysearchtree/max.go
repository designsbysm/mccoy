package binarysearchtree

func (t *Tree[T]) Max() *T {
	t.lock.RLock()
	defer t.lock.RUnlock()

	node := t.root
	if node == nil {
		return nil
	}

	for {
		if node.right == nil {
			return &node.value
		}

		node = node.right
	}
}
