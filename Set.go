package Set

type Set[T comparable] struct {
	items map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		items: make(map[T]struct{}),
	}
}

func (s *Set[T]) Add(item T) {
	s.items[item] = struct{}{}
}

func (s *Set[T]) Remove(item T) {
	delete(s.items, item)
}

func (s *Set[T]) Contains(item T) bool {
	_, exist := s.items[item]
	return exist
}

func (s *Set[T]) Size() int {
	return len(s.items)
}

func (s *Set[T]) Clear() {
	s.items = make(map[T]struct{})
}
