package storage

import(
    "fmt"
)

type InMemoryStorage struct {
    data map[string][]byte
}

func NewInMemoryStorage() *InMemoryStorage {
    return &InMemoryStorage{
        data: make(map[string][]byte),
    }
}

func (s *InMemoryStorage) Get(key string) ([]byte, error) {
    val, ok := s.data[key]
    if !ok {
        return nil, fmt.Errorf("get %q: %w", key, ErrNotFound)
    }
    return val, nil
}

func (s *InMemoryStorage) Set(key string, value []byte) error {
    s.data[key] = value
    return nil
}

func (s *InMemoryStorage) Delete(key string) error {
    if _, ok := s.data[key]; !ok {
        return fmt.Errorf("delete %q: %w", key, ErrNotFound)
    }
    delete(s.data, key)
    return nil
}