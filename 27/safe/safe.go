package safe

import (
	"fmt"
	"sync"
)

type Safe struct {
	mu      sync.Mutex
	counter map[string]int
}

func New() *Safe {

	my := map[string]int{
		"a": 0,
		"b": 0,
	}
	return &Safe{counter: my}
}

func (s *Safe) Increment(key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.counter[key]
	if !ok {
		return fmt.Errorf("key %s not found", key)
	}

	s.counter[key]++
	return nil
}

func (s *Safe) IncrementByValue(key string, number int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.counter[key]
	if !ok {
		return fmt.Errorf("key %s not found", key)
	}

	s.counter[key] += number
	return nil
}

func (s *Safe) Get(key string) int {
	_, ok := s.counter[key]
	if !ok {
		return 0
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.counter[key]
}
