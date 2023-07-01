package set

import "github.com/designsbysm/mccoy"

func (s *T) Delete(i mccoy.Item) {
	delete(s.items, i)
}
