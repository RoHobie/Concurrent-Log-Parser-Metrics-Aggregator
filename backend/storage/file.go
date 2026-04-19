package storage

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type FileStorage struct {
	baseDir string
}

func NewFileStorage(baseDir string) (*FileStorage, error) {
    if err := os.MkdirAll(baseDir, 0755); err != nil {
        return nil, fmt.Errorf("init storage: %w", err)
    }
    return &FileStorage{baseDir: baseDir}, nil
}

func (s *FileStorage) path(key string) string {
	return filepath.Join(s.baseDir, key)
}

func (s *FileStorage) Get(key string) ([]byte, error) {
    data, err := os.ReadFile(s.path(key))
    if err != nil {
        if errors.Is(err, os.ErrNotExist) {
            return nil, fmt.Errorf("get %q: %w", key, ErrNotFound)
        }
        return nil, fmt.Errorf("get %q: %w", key, err)
    }
    return data, nil
}

func (s *FileStorage) Set(key string, value []byte) error {
	if err := os.WriteFile(s.path(key), value, 0644); err != nil {
		return fmt.Errorf("set %q: %w", key, err)
	}
	return nil
}

func (s *FileStorage) Delete(key string) error {
	err := os.Remove(s.path(key))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("delete %q: %w", key, ErrNotFound)
		}
		return fmt.Errorf("delete %q: %w", key, err)
	}
	return nil
}
