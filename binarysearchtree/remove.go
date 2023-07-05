package binarysearchtree

func (t *Tree[T]) Remove(key int) {
	t.lock.Lock()
	defer t.lock.Unlock()

	removeNode(t.root, key)
}

func removeNode[T any](node *Node[T], key int) *Node[T] {
	if node == nil {
		return nil
	} else if key < node.key {
		node.left = removeNode(node.left, key)
		return node
	} else if key > node.key {
		node.right = removeNode(node.right, key)
		return node
	} else if node.left == nil && node.right == nil {
		node = nil
		return nil
	} else if node.left == nil {
		node = node.right
		return node
	} else if node.right == nil {
		node = node.left
		return node
	}

	leftmostrightNode := node.right
	for {
		if leftmostrightNode != nil && leftmostrightNode.left != nil {
			leftmostrightNode = leftmostrightNode.left
		} else {
			break
		}
	}

	node.key, node.value = leftmostrightNode.key, leftmostrightNode.value
	node.right = removeNode(node.right, node.key)

	return node
}
