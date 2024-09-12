package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
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

func measure(allFiles [][]byte, pattern string) {
	// Compile regex
	r, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal(err)
	}

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

			matches := r.FindAllIndex(input, -1)

			mu.Lock()
			count += len(matches)
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
	measure(allFiles, `[\w\.+-]+@[\w\.-]+\.[\w\.-]+`)

	// URI
	measure(allFiles, `[\w]+://[^/\s?#]+[^\s?#]+(?:\?[^\s#]*)?(?:#[^\s]*)?`)

	// IP
	measure(allFiles, `(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9])\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9])`)

	// Long date pattern
	day := `[0-3]?[0-9]`
	month := `` +
		`January?|February?|March|A[pv]ril|Ma[iy]|Jun[ei]|Jul[iy]|August|September|O[ck]tober|November|De[csz]ember|` +
		`Jan|Feb|M[aä]r|Apr|Jun|Jul|Aug|Sep|O[ck]t|Nov|De[cz]|` +
		`Januari|Februari|Maret|Mei|Agustus|` +
		`Jänner|Feber|März|` +
		`janvier|février|mars|juin|juillet|aout|septembre|octobre|novembre|décembre|` +
		`Ocak|Şubat|Mart|Nisan|Mayıs|Haziran|Temmuz|Ağustos|Eylül|Ekim|Kasım|Aralık|` +
		`Oca|Şub|Mar|Nis|Haz|Tem|Ağu|Eyl|Eki|Kas|Ara`
	year := `199[0-9]|20[0-3][0-9]`
	longDatePattern := fmt.Sprintf(`(?i)`+
		`(%[2]s)\s(%[3]s)(?:st|nd|rd|th)?,?\s(%[1]s)`+
		`|`+
		`(%[3]s)(?:st|nd|rd|th|\.)?\s(?:of\s)?(%[2]s)[,.]?\s(%[1]s)`,
		year, month, day)
	measure(allFiles, longDatePattern)
}
