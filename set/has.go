package set

import "github.com/designsbysm/mccoy"

func (s *T) Has(i mccoy.Item) bool {
	_, ok := s.items[i]
	return ok
}
