package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	regexp "github.com/grafana/regexp"
)

func measure(data []byte, pattern string) {
	r, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	matches := r.FindAllIndex(data, -1)
	count := len(matches)
	elapsed := time.Since(start)

	fmt.Printf("%f - %v\n", float64(elapsed)/float64(time.Millisecond), count)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: benchmark <filename>")
		os.Exit(1)
	}

	filerc, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer filerc.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(filerc)
	data := buf.Bytes()

	// Email
	measure(data, `[\w\.+-]+@[\w\.-]+\.[\w\.-]+`)

	// URI
	measure(data, `[\w]+://[^/\s?#]+[^\s?#]+(?:\?[^\s#]*)?(?:#[^\s]*)?`)

	// IP
	measure(data, `(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9])\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9])`)

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
	measure(data, longDatePattern)
}
