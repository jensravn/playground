package frequency

import (
	"bufio"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"
)

func TermConcurrent(text string, n int, stopWords string) []WordFreq {
	// Extract words
	words := getWordsConcurrent(text)

	// Load stop words
	stopWordMap := getStopWordMap(stopWords)

	// Count word frequencies
	wordFreqs := getWordFreqsConcurrent(words, stopWordMap)

	// Sort desc freq
	sort.Slice(wordFreqs, func(a, b int) bool {
		return wordFreqs[a].Freq > wordFreqs[b].Freq
	})

	// Prepare result of n elements
	topN := wordFreqs[:n]

	return topN
}

func getWordsConcurrent(text string) []string {
	fmt.Printf("0 START\n")
	start := time.Now()

	lines := make(chan string)
	wordChunks := make(chan []string)

	// Send tasks to channel
	var wg1 sync.WaitGroup
	wg1.Add(1)
	go func() {
		defer wg1.Done()
		lowerText := strings.ToLower(text)
		scanner := bufio.NewScanner(strings.NewReader(lowerText))
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			lines <- scanner.Text()
		}
	}()
	go func() {
		wg1.Wait()
		close(lines)
		fmt.Printf("1 lines closed - %s\n", time.Since(start))
	}()

	// Concurrent workers
	regex := regexp.MustCompile(`[\W_']+`)
	var wg2 sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			for line := range lines {
				cleanedText := regex.ReplaceAllString(line, " ")
				var words []string
				for _, word := range strings.Split(cleanedText, " ") {
					words = append(words, strings.ToLower(word))
				}
				wordChunks <- words
			}
		}()
	}
	go func() {
		wg2.Wait()
		close(wordChunks)
		fmt.Printf("2 wordChunks closed - %s\n", time.Since(start))
	}()

	// Aggregate worker results
	var words []string
	var wg3 sync.WaitGroup
	wg3.Add(1)
	go func() {
		defer wg3.Done()
		for wordChunk := range wordChunks {
			words = append(words, wordChunk...)
		}
	}()
	wg3.Wait()
	fmt.Printf("3 Words Split - %s\n", time.Since(start))
	return words
}

func getWordFreqsConcurrent(words []string, stopWordMap map[string]bool) []WordFreq {
	tasks := make(chan []string)
	results := make(chan *map[string]int)

	// Send tasks to channel
	var wg1 sync.WaitGroup
	wg1.Add(1)
	go func() {
		defer wg1.Done()
		size := 1000000
		var j int
		for i := 0; i < len(words); i += size {
			j += size
			if j > len(words) {
				j = len(words)
			}
			tasks <- words[i:j]
		}
	}()
	go func() {
		wg1.Wait()
		close(tasks)
	}()

	// Concurrent workers
	var wg2 sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			for t := range tasks {
				results <- getWordFreqMap(t, stopWordMap)
			}
		}()
	}
	go func() {
		wg2.Wait()
		close(results)
	}()

	// Receive results
	wordFreqMap := map[string]int{}
	var wg3 sync.WaitGroup
	wg3.Add(1)
	go func() {
		defer wg3.Done()
		for wordFreqChunk := range results {
			for word, freq := range *wordFreqChunk {
				wordFreqMap[word] += freq
			}
		}
	}()
	wg3.Wait()
	wordFreqs := []WordFreq{}
	for word, freq := range wordFreqMap {
		wordFreqs = append(wordFreqs, WordFreq{word, freq})
	}
	return wordFreqs
}

func getWordFreqMap(words []string, stopWordMap map[string]bool) *map[string]int {
	wordFreqMap := map[string]int{}
	for _, word := range words {
		if !stopWordMap[word] {
			if _, ok := wordFreqMap[word]; ok {
				wordFreqMap[word]++
			} else {
				wordFreqMap[word] = 1
			}
		}
	}
	return &wordFreqMap
}
