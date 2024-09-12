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
	"unicode/utf8"

	regexp "github.com/dlclark/regexp2"
	"golang.org/x/sync/semaphore"
)

var (
	rxEmail    = regexp.MustCompile(`[\w\.+-]+@[\w\.-]+\.[\w\.-]+`, 0)
	rxURI      = regexp.MustCompile(`[\w]+://[^/\s?#]+[^\s?#]+(?:\?[^\s#]*)?(?:#[^\s]*)?`, 0)
	rxIP       = regexp.MustCompile(`(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9])\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9])`, 0)
	rxLongDate = regexp.MustCompile(`(?i)(January?|February?|March|A[pv]ril|Ma[iy]|Jun[ei]|Jul[iy]|August|September|O[ck]tober|November|De[csz]ember|Jan|Feb|M[aä]r|Apr|Jun|Jul|Aug|Sep|O[ck]t|Nov|De[cz]|Januari|Februari|Maret|Mei|Agustus|Jänner|Feber|März|janvier|février|mars|juin|juillet|aout|septembre|octobre|novembre|décembre|Ocak|Şubat|Mart|Nisan|Mayıs|Haziran|Temmuz|Ağustos|Eylül|Ekim|Kasım|Aralık|Oca|Şub|Mar|Nis|Haz|Tem|Ağu|Eyl|Eki|Kas|Ara)\s([0-3]?[0-9])(?:st|nd|rd|th)?,?\s(199[0-9]|20[0-3][0-9])|([0-3]?[0-9])(?:st|nd|rd|th|\.)?\s(?:of\s)?(January?|February?|March|A[pv]ril|Ma[iy]|Jun[ei]|Jul[iy]|August|September|O[ck]tober|November|De[csz]ember|Jan|Feb|M[aä]r|Apr|Jun|Jul|Aug|Sep|O[ck]t|Nov|De[cz]|Januari|Februari|Maret|Mei|Agustus|Jänner|Feber|März|janvier|février|mars|juin|juillet|aout|septembre|octobre|novembre|décembre|Ocak|Şubat|Mart|Nisan|Mayıs|Haziran|Temmuz|Ağustos|Eylül|Ekim|Kasım|Aralık|Oca|Şub|Mar|Nis|Haz|Tem|Ağu|Eyl|Eki|Kas|Ara)[,.]?\s(199[0-9]|20[0-3][0-9])`, 0)
)

func bytes2runes(src []byte) []rune {
	runes := make([]rune, 0, len(src))
	for len(src) > 0 {
		r, size := utf8.DecodeRune(src)
		runes = append(runes, r)
		src = src[size:]
	}
	return runes
}

func openAllFiles(dir string) ([][]rune, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var files [][]rune
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		srcPath := filepath.Join(dir, entry.Name())
		file, err := os.ReadFile(srcPath)
		if err != nil {
			return nil, err
		}

		files = append(files, bytes2runes(file))
	}

	return files, nil
}

func measure(allFiles [][]rune, r *regexp.Regexp) {
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
		go func(input []rune) {
			defer func() {
				wg.Done()
				sem.Release(1)
			}()

			var nMatches int
			m, _ := r.FindRunesMatch(allFiles[i])
			for m != nil {
				nMatches++
				m, _ = r.FindNextMatch(m)
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
	measure(allFiles, rxEmail)

	// URI
	measure(allFiles, rxURI)

	// IP
	measure(allFiles, rxIP)

	// Long date pattern
	measure(allFiles, rxLongDate)
}
