package linkedlist

import (
	"github.com/designsbysm/mccoy"
)

// IDEA: make the list heterogeneous with generics
type T struct {
	head mccoy.Item
	tail *T
}
