package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"Concurrent-Log-Parser-Metrics-Aggregator/backend/storage"
)

func main() {
	// ensure testdata directory exists
	baseDir := "./testdata"
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		log.Fatalf("failed to create testdata dir: %v", err)
	}

	// init storages
	fileStore, err := storage.NewFileStorage(baseDir)
	if err != nil {
		log.Fatalf("failed to init file storage: %v", err)
	}

	memStore := storage.NewInMemoryStorage()

	fmt.Println("---- Testing FileStorage ----")
	testStorage(fileStore)

	fmt.Println("\n---- Testing InMemoryStorage ----")
	testStorage(memStore)
}

func testStorage(s storage.Storage) {
	key := "example.txt"
	value := []byte("hello world")

	// 1. Ensure key does NOT exist
	_, err := s.Get(key)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			fmt.Println("OK: key not found initially")
		} else {
			fmt.Printf("Unexpected error: %v\n", err)
		}
	}

	// 2. Set value
	if err := s.Set(key, value); err != nil {
		fmt.Printf("Set failed: %v\n", err)
		return
	}
	fmt.Println("OK: Set successful")

	// 3. Get value
	data, err := s.Get(key)
	if err != nil {
		fmt.Printf("Get failed: %v\n", err)
		return
	}
	fmt.Printf("OK: Get successful → %s\n", string(data))

	// 4. Delete key
	if err := s.Delete(key); err != nil {
		fmt.Printf("Delete failed: %v\n", err)
		return
	}
	fmt.Println("OK: Delete successful")

	// 5. Confirm deletion
	_, err = s.Get(key)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			fmt.Println("OK: key not found after delete")
		} else {
			fmt.Printf("Unexpected error after delete: %v\n", err)
		}
	}
}