package main

func main() {
	s := newStore()
	s.insert("a", "1")
	s.update("a", "2")
	s.upsert("b", "3")
	println(s.contains("a"))
	s.delete("a")
	println(s.contains("a"))
}

type store struct {
	m map[string]string
}

func newStore() *store {
	return &store{
		m: map[string]string{},
	}
}

func (s *store) insert(k, v string) {
	if _, ok := s.m[k]; !ok {
		s.m[k] = v
	}
}

func (s *store) update(k, v string) {
	if _, ok := s.m[k]; ok {
		s.m[k] = v
	}
}

func (s *store) upsert(k, v string) {
	s.m[k] = v
}

func (s *store) contains(k string) bool {
	_, ok := s.m[k]
	return ok
}

func (s *store) delete(k string) {
	delete(s.m, k)
}
