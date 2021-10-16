package mccoy

type Set struct {
	items map[Item]bool
}

func (s *Set) Add(i Item) {
	if s.items == nil {
		s.items = make(map[Item]bool)
	}

	if _, ok := s.items[i]; !ok {
		s.items[i] = true
	}
}

func (s *Set) Clear() {
	s.items = make(map[Item]bool)
}

func (s *Set) Delete(i Item) {
	delete(s.items, i)
}

func (s *Set) Has(i Item) bool {
	_, ok := s.items[i]
	return ok
}

func (s Set) Items() []Item {
	items := []Item{}

	for i := range s.items {
		items = append(items, i)
	}

	return items
}

func (s Set) Size() int {
	return len(s.items)
}
