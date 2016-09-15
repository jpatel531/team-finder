package fetcher

import "sync"

type intSet struct {
	*sync.RWMutex
	data map[int]bool
}

func newIntSet() *intSet {
	return &intSet{
		&sync.RWMutex{},
		make(map[int]bool),
	}
}

func (s *intSet) add(i int) {
	s.Lock()
	defer s.Unlock()
	s.data[i] = true
}

func (s *intSet) exists(i int) bool {
	s.RLock()
	defer s.RUnlock()
	return s.data[i]
}
