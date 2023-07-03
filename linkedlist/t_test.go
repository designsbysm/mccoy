package linkedlist

import "fmt"

func createList() *Node[int] {
	list := New(1)

	for i := 2; i < 10; i++ {
		list = list.Cons(i)
	}

	return list.Reverse()
}

func printList[T any](l *Node[T]) {
	fmt.Print(" ")

	if l == nil {
		fmt.Println("<empty>")
		return
	}

	for node := l; node != nil; node = node.tail {
		fmt.Print(node.head, " ")
	}

	fmt.Print("\n")
}
