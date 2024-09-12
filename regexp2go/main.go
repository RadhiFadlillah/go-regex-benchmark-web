package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp2go/internal/re"
	"runtime"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

type RegexPattern[MatchType any] interface {
	Find(s []byte) (matches MatchType, pos int, ok bool)
}

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

func measure[T [1][]byte | [7][]byte](allFiles [][]byte, pattern RegexPattern[T]) {
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
			var cursor int
			var nMatches int
			limit := len(input)

			for {
				matches, pos, ok := pattern.Find(input[cursor:])
				if ok {
					nMatches++
					cursor += pos + len(matches[0])
					if cursor <= limit {
						continue
					}
				}
				break
			}

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

	// Email
	measure(allFiles, re.RxEmail{})

	// URI
	measure(allFiles, re.RxURI{})

	// IP
	measure(allFiles, re.RxIP{})

	// Long Date
	measure(allFiles, re.RxLongDate{})
}
