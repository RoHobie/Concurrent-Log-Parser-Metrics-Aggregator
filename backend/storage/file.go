package storage

import(
	"os"
	"path/filepath"
)

type FileStorage struct {
    baseDir string
}

func NewFileStorage(baseDir string) *FileStorage {
    os.MkdirAll(baseDir, 0755)
    return &FileStorage{baseDir: baseDir}
}

func (s *FileStorage) path(key string) string {
    return filepath.Join(s.baseDir, key)
}

func (s *FileStorage) Get(key string) ([]byte, error) {
    return os.ReadFile(s.path(key))
}

func (s *FileStorage) Set(key string, value []byte) error {
    return os.WriteFile(s.path(key), value, 0644)
}

func (s *FileStorage) Delete(key string) error {
    return os.Remove(s.path(key))
}