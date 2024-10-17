package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

var exams = []struct {
	Exam      string
	Questions int
	PageSize  int
}{
	{Exam: "associate-cloud-engineer", Questions: 266},
	{Exam: "cloud-digital-leader", Questions: 70},
	{Exam: "professional-cloud-architect", Questions: 276},
	{Exam: "professional-cloud-database-engineer", Questions: 132, PageSize: 4},
	{Exam: "professional-cloud-developer", Questions: 279},
	{Exam: "professional-cloud-devops-engineer", Questions: 166, PageSize: 4},
	{Exam: "professional-cloud-network-engineer", Questions: 172, PageSize: 4},
	{Exam: "professional-cloud-security-engineer", Questions: 244},
	{Exam: "professional-data-engineer", Questions: 311, PageSize: 10},
	{Exam: "professional-google-workspace-administrator", Questions: 132, PageSize: 4},
	{Exam: "professional-machine-learning-engineer", Questions: 285},
}

func main() {
	t := time.Now()
	seed := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC).UnixNano()
	r := rand.New(rand.NewPCG(uint64(seed), 2))
	n := r.IntN(len(exams))
	e := exams[n]
	q := r.IntN(e.Questions) + 1
	page := getPage(q, e.PageSize)
	url := "https://www.examtopics.com/exams/google/%s/view/%d\n"
	fmt.Printf("%s #%d\n", e.Exam, q)
	fmt.Printf(url, e.Exam, page)
}

func getPage(question int, pageSize int) int {
	if pageSize == 0 {
		pageSize = 5
	}
	return (question-1)/pageSize + 1
}
