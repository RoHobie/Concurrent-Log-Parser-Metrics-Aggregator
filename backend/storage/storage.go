package storage

import "errors"

var ErrNotFound = errors.New("not found")

type Storage interface {
    Get(key string) ([]byte, error)
    Set(key string, value []byte) error
    Delete(key string) error
}

