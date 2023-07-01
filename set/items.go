package set

import "github.com/designsbysm/mccoy"

func (s T) Items() []mccoy.Item {
	items := []mccoy.Item{}

	for i := range s.items {
		items = append(items, i)
	}

	return items
}
