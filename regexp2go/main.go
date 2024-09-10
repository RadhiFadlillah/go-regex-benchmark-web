package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp2/internal/re"
	"time"
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
	var count int
	start := time.Now()

	for _, data := range allFiles {
		var cursor int
		limit := len(data)

		for {
			matches, pos, ok := pattern.Find(data[cursor:])
			if ok {
				count++
				cursor += pos + len(matches[0])
				if cursor <= limit {
					continue
				}
			}
			break
		}
	}

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
