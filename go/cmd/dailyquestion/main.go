package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := rand.Intn(195) + 1
	fmt.Printf("https://www.examtopics.com/exams/google/professional-cloud-architect/view/%d\n", (n/5)+1)
	fmt.Printf("Question #%d\n", n)
}
