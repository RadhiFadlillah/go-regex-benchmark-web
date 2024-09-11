package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
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
	var count int
	start := time.Now()
	for i := range allFiles {
		nMatches := rxFinder(allFiles[i])
		count += nMatches
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

	measure(allFiles, findEmails)
	measure(allFiles, findURIs)
	measure(allFiles, findIPs)
	measure(allFiles, findLongDatePattern)
}
