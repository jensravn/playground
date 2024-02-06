package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := rand.Intn(195) + 1
	page := getPage(n)
	url := "https://www.examtopics.com/exams/google/professional-cloud-architect/view/%d\n"
	fmt.Printf(url, page)
	fmt.Printf("Question #%d\n", n)
}

func getPage(question int) int {
	return (question-1)/5 + 1
}
