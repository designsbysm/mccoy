package binarysearchtree

func (t *Tree[T]) Dequeue() *Node[T] {
	if t.root == nil {
		return nil
	}

	t.lock.RLock()
	defer t.lock.RUnlock()

	min := t.Min()
	t.Remove(min.key)

	min.left = nil
	min.right = nil

	return min
}
