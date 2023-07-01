package set

import "github.com/designsbysm/mccoy"

func (s *T) Add(i mccoy.Item) {
	if s.items == nil {
		s.Clear()
	}

	if _, ok := s.items[i]; !ok {
		s.items[i] = true
	}
}
