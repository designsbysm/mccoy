package linkedlist

import (
	"fmt"
)

func createList() *T {
	list := New(1)

	for i := 2; i < 10; i++ {
		list = list.Prepend(i)
	}

	list = list.Reverse()

	return list
}

func printList(list *T) {
	for {
		if list == nil {
			break
		} else if list.head == nil {
			fmt.Println("<empty>")
			return
		}

		fmt.Print(list.head, " ")
		list = list.tail
	}
	fmt.Print("\n")
}
