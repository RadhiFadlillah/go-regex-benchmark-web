package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp2/internal/re"
	"time"
)

type RegexPattern[MatchType any] interface {
	Find(s []byte) (matches MatchType, pos int, ok bool)
}

func measure[T [1][]byte | [7][]byte](data []byte, pattern RegexPattern[T]) {
	var count int
	var cursor int
	limit := len(data)

	start := time.Now()
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
	measure(data, re.RxEmail{})

	// URI
	measure(data, re.RxURI{})

	// IP
	measure(data, re.RxIP{})

	// Long Date
	measure(data, re.RxLongDate{})
}
