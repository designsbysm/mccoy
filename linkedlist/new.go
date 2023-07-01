package linkedlist

import "github.com/designsbysm/mccoy"

func New(i mccoy.Item) *T {
	return &T{
		i,
		nil,
	}
}

// func New() *T {
// 	return &T{
// 		nil,
// 		nil,
// 	}
// }
