package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"
)

func measure(data []rune, rxFinder func([]rune) int) {
	start := time.Now()
	count := rxFinder(data)
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
	data := []rune(buf.String())

	measure(data, findEmails)
	measure(data, findURIs)
	measure(data, findIPs)
	measure(data, findLongDatePattern)
}
