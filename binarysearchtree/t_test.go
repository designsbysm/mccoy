package binarysearchtree

func createTree() *Tree[int] {
	tree := New[int]()
	tree.Enqueue(8, 8)
	tree.Enqueue(3, 3)
	tree.Enqueue(-1, -1)
	tree.Enqueue(10, 10)
	tree.Enqueue(1, 1)
	tree.Enqueue(0, 0)
	tree.Enqueue(6, 6)
	tree.Enqueue(2, 2)
	tree.Enqueue(11, 11)

	return tree
}
