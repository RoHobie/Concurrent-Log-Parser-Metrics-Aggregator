package storage

import(
    "fmt"
    "sync"
)

type InMemoryStorage struct {
	data map[string][]byte
	mu   sync.RWMutex
}

func (s *InMemoryStorage) Get(key string) ([]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	val, ok := s.data[key]
	if !ok {
		return nil, fmt.Errorf("get %q: %w", key, ErrNotFound)
	}
	return val, nil
}

func (s *InMemoryStorage) Set(key string, value []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value
	return nil
}

func (s *InMemoryStorage) Delete(key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.data[key]; !ok {
		return fmt.Errorf("delete %q: %w", key, ErrNotFound)
	}
	delete(s.data, key)
	return nil
}