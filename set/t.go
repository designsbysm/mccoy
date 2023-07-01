package set

import "github.com/designsbysm/mccoy"

// IDEA: make the list heterogeneous with generics
type T struct {
	items map[mccoy.Item]bool
}
