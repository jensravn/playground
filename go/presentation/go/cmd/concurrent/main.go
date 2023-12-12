package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jensravn/term-frequency/go/lib/frequency"
)

func main() {
	// Read input files
	text, err := os.ReadFile("../../../_resources/pride-and-prejudice500.txt")
	if err != nil {
		log.Fatal(err)
	}
	n := 25
	stopwords, err := os.ReadFile("../../../_resources/stop_words.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("START\n")
	start := time.Now()

	// Term frequency
	topN := frequency.TermConcurrent(string(text), n, string(stopwords))

	fmt.Printf("END - %s\n\n", time.Since(start).Round(time.Millisecond))

	// Print top n words
	for _, x := range topN {
		fmt.Printf("%s - %d\n", x.Word, x.Freq)
	}
}
