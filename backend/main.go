package main

import (
	"fmt"
	"log"
	"Concurrent-Log-Parser-Metrics-Aggregator/backend/storage" 
)

func testStorage(name string, s storage.Storage) {
	fmt.Printf("\n--- Testing %s ---\n", name)

	// Set
	if err := s.Set("hello", []byte("world")); err != nil {
		log.Fatalf("Set failed: %v", err)
	}
	fmt.Println("Set 'hello' = 'world'")

	// Get existing key
	val, err := s.Get("hello")
	if err != nil {
		log.Fatalf("Get failed: %v", err)
	}
	fmt.Printf("Get 'hello' = '%s'\n", val)

	// Get missing key
	_, err = s.Get("missing")
	if err == storage.ErrNotFound {
		fmt.Println("Get 'missing' correctly returned ErrNotFound")
	} else {
		log.Fatalf("Expected ErrNotFound, got: %v", err)
	}

	// Delete
	if err := s.Delete("hello"); err != nil {
		log.Fatalf("Delete failed: %v", err)
	}
	fmt.Println("Deleted 'hello'")

	// Get after delete
	_, err = s.Get("hello")
	if err == storage.ErrNotFound {
		fmt.Println("Get after delete correctly returned ErrNotFound")
	} else {
		log.Fatalf("Expected ErrNotFound after delete, got: %v", err)
	}
}

func main() {
	// Test in-memory storage
	mem := storage.NewInMemoryStorage()
	testStorage("InMemoryStorage", mem)

	// Test file storage
	file := storage.NewFileStorage("./testdata")
	testStorage("FileStorage", file)

	fmt.Println("\nAll tests passed!")
}