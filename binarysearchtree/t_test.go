package binarysearchtree

func createTree() *Tree[int] {
	tree := New[int]()
	tree.Insert(8, 8)
	tree.Insert(3, 3)
	tree.Insert(-1, -1)
	tree.Insert(10, 10)
	tree.Insert(1, 1)
	tree.Insert(0, 0)
	tree.Insert(6, 6)
	tree.Insert(2, 2)
	tree.Insert(11, 11)

	return tree
}
