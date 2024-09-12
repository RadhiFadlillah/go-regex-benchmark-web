package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
	"unicode/utf8"

	regexp "github.com/dlclark/regexp2"
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
	var count int
	start := time.Now()
	for i := range allFiles {
		m, _ := r.FindRunesMatch(allFiles[i])
		for m != nil {
			count++
			m, _ = r.FindNextMatch(m)
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
	measure(allFiles, rxEmail)

	// URI
	measure(allFiles, rxURI)

	// IP
	measure(allFiles, rxIP)

	// Long date pattern
	measure(allFiles, rxLongDate)
}
