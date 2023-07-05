package binarysearchtree

func (t *Tree[T]) Enqueue(key int, value T) {
	t.lock.Lock()
	defer t.lock.Unlock()

	node := &Node[T]{
		key:   key,
		value: value,
		left:  nil,
		right: nil,
	}

	if t.root == nil {
		t.root = node
	} else {
		insertNode(t.root, node)
	}
}

func insertNode[T any](root *Node[T], new *Node[T]) {
	if new.key < root.key {
		if root.left == nil {
			root.left = new
		} else {
			insertNode(root.left, new)
		}
	} else {
		if root.right == nil {
			root.right = new
		} else {
			insertNode(root.right, new)
		}
	}
}
