package linkedlist

import (
	"fmt"
	"strconv"
)

func createList() *Node[int] {
	list := New(1)

	for i := 2; i < 10; i++ {
		list = list.Cons(i)
	}

	return list.Reverse()
}

func printList(l *Node[int]) {
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

func toString(l *Node[int]) string {
	var str string

	l.ForEach(func(node *Node[int]) {
		str += strconv.Itoa(node.head)
	})

	return str
}
