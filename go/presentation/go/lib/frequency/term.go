package frequency

import (
	"regexp"
	"sort"
	"strings"
)

type WordFreq struct {
	Word string
	Freq int
}

func Term(text string, n int, stopWords string) []WordFreq {
	// Extract words
	words := getWords(text)

	// Load stop words
	stopWordMap := getStopWordMap(stopWords)

	// Count word frequencies
	wordFreqs := getWordFreqs(words, stopWordMap)

	// Sort desc freq
	sort.Slice(wordFreqs, func(a, b int) bool {
		return wordFreqs[a].Freq > wordFreqs[b].Freq
	})

	// Prepare result of n elements
	topN := wordFreqs[:n]

	return topN
}

func getWords(text string) []string {
	words := []string{}
	re := regexp.MustCompile(`[\W_']+`)
	cleanedText := re.ReplaceAllString(strings.ToLower(text), " ")
	for _, word := range strings.Split(cleanedText, " ") {
		words = append(words, strings.ToLower(word))
	}
	return words
}

func getStopWordMap(stopWords string) map[string]bool {
	stopWordMap := map[string]bool{}
	for _, word := range strings.Split(stopWords, ",") {
		stopWordMap[word] = true
	}
	return stopWordMap
}

func getTopN(wordFreqs []WordFreq, n int) map[string]int {
	topN := map[string]int{}
	for i := 0; i < n; i++ {
		topN[wordFreqs[i].Word] = wordFreqs[i].Freq
	}
	return topN
}

func getWordFreqs(words []string, stopWordMap map[string]bool) []WordFreq {
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
	wordFreqs := []WordFreq{}
	for word, freq := range wordFreqMap {
		wordFreqs = append(wordFreqs, WordFreq{word, freq})
	}
	return wordFreqs
}
