package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

func openAllFiles(dir string) ([][]byte, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var files [][]byte
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		srcPath := filepath.Join(dir, entry.Name())
		file, err := os.ReadFile(srcPath)
		if err != nil {
			return nil, err
		}

		files = append(files, file)
	}

	return files, nil
}

func measure(allFiles [][]byte, rxFinder func([]byte) int) {
	// Prepare counter
	var mu sync.Mutex
	var count int

	// Prepare wait group
	var wg sync.WaitGroup
	ctx := context.TODO()
	maxWorkers := runtime.GOMAXPROCS(0)
	sem := semaphore.NewWeighted(int64(maxWorkers))

	start := time.Now()
	for i := range allFiles {
		if err := sem.Acquire(ctx, 1); err != nil {
			log.Fatalf("failed to acquire semaphore: %v", err)
		}

		wg.Add(1)
		go func(input []byte) {
			defer func() {
				wg.Done()
				sem.Release(1)
			}()

			nMatches := rxFinder(allFiles[i])

			mu.Lock()
			count += nMatches
			mu.Unlock()
		}(allFiles[i])
	}

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("%f - %v\n", float64(elapsed)/float64(time.Millisecond), count)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: benchmark <filedir>")
		os.Exit(1)
	}

	allFiles, err := openAllFiles(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	measure(allFiles, findEmails)
	measure(allFiles, findURIs)
	measure(allFiles, findIPs)
	measure(allFiles, findLongDatePattern)
}
