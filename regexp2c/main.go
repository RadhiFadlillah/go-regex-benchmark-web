package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	regexp "github.com/dlclark/regexp2"
)

var (
	rxEmail    = regexp.MustCompile(`[\w\.+-]+@[\w\.-]+\.[\w\.-]+`, 0)
	rxURI      = regexp.MustCompile(`[\w]+://[^/\s?#]+[^\s?#]+(?:\?[^\s#]*)?(?:#[^\s]*)?`, 0)
	rxIP       = regexp.MustCompile(`(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9])\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9])`, 0)
	rxLongDate = regexp.MustCompile(`(?i)(January?|February?|March|A[pv]ril|Ma[iy]|Jun[ei]|Jul[iy]|August|September|O[ck]tober|November|De[csz]ember|Jan|Feb|M[aä]r|Apr|Jun|Jul|Aug|Sep|O[ck]t|Nov|De[cz]|Januari|Februari|Maret|Mei|Agustus|Jänner|Feber|März|janvier|février|mars|juin|juillet|aout|septembre|octobre|novembre|décembre|Ocak|Şubat|Mart|Nisan|Mayıs|Haziran|Temmuz|Ağustos|Eylül|Ekim|Kasım|Aralık|Oca|Şub|Mar|Nis|Haz|Tem|Ağu|Eyl|Eki|Kas|Ara)\s([0-3]?[0-9])(?:st|nd|rd|th)?,?\s(199[0-9]|20[0-3][0-9])|([0-3]?[0-9])(?:st|nd|rd|th|\.)?\s(?:of\s)?(January?|February?|March|A[pv]ril|Ma[iy]|Jun[ei]|Jul[iy]|August|September|O[ck]tober|November|De[csz]ember|Jan|Feb|M[aä]r|Apr|Jun|Jul|Aug|Sep|O[ck]t|Nov|De[cz]|Januari|Februari|Maret|Mei|Agustus|Jänner|Feber|März|janvier|février|mars|juin|juillet|aout|septembre|octobre|novembre|décembre|Ocak|Şubat|Mart|Nisan|Mayıs|Haziran|Temmuz|Ağustos|Eylül|Ekim|Kasım|Aralık|Oca|Şub|Mar|Nis|Haz|Tem|Ağu|Eyl|Eki|Kas|Ara)[,.]?\s(199[0-9]|20[0-3][0-9])`, 0)
)

func measure(data []rune, r *regexp.Regexp) {
	start := time.Now()

	var matches [][]rune
	m, _ := r.FindRunesMatch(data)
	for m != nil {
		matches = append(matches, m.Runes())
		m, _ = r.FindNextMatch(m)
	}

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
	data := bytes.Runes(buf.Bytes())

	// Email
	measure(data, rxEmail)

	// URI
	measure(data, rxURI)

	// IP
	measure(data, rxIP)

	// Long date pattern
	measure(data, rxLongDate)
}
