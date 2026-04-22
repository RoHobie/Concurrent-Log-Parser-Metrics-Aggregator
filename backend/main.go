package main

import (
	"fmt"
	"log"
	"os"

	"Concurrent-Log-Parser-Metrics-Aggregator/backend/processor"
	"Concurrent-Log-Parser-Metrics-Aggregator/backend/storage"
)

func main() {
	baseDir := "./testdata"
	_ = os.MkdirAll(baseDir, 0755)

	store, err := storage.NewFileStorage(baseDir)
	if err != nil {
		log.Fatal(err)
	}

	wp := processor.NewWorkerPool(store, 10)

	for i := 0; i < 1000; i++ {
		wp.Submit(processor.Job{
			Key:  fmt.Sprintf("log-%d.txt", i),
			Data: []byte(fmt.Sprintf("log line %d", i)),
		})
	}

	wp.Close()

	fmt.Println("Processed:", wp.Count())
}