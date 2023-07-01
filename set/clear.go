package set

import "github.com/designsbysm/mccoy"

func (s *T) Clear() {
	s.items = make(map[mccoy.Item]bool)
}
