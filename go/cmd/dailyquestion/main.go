package main

import (
	"fmt"
	"math/rand"
)

var exams = []struct {
	Exam      string
	Questions int
	PageSize  int
}{
	{Exam: "associate-cloud-engineer", Questions: 266},
	{Exam: "cloud-digital-leader", Questions: 286},
	{Exam: "professional-cloud-architect", Questions: 276},
	{Exam: "professional-cloud-database-engineer", Questions: 132},
	{Exam: "professional-cloud-developer", Questions: 279},
	{Exam: "professional-cloud-devops-engineer", Questions: 166},
	{Exam: "professional-cloud-network-engineer", Questions: 172, PageSize: 4},
	{Exam: "professional-cloud-security-engineer", Questions: 244},
	{Exam: "professional-collaboration-engineer", Questions: 79},
	{Exam: "professional-data-engineer", Questions: 311},
	{Exam: "professional-google-workspace-administrator", Questions: 132},
	{Exam: "professional-machine-learning-engineer", Questions: 285},
}

func main() {
	n := rand.Intn(len(exams))
	e := exams[n]
	q := rand.Intn(e.Questions) + 1
	page := getPage(q, e.PageSize)
	url := "https://www.examtopics.com/exams/google/%s/view/%d\n"
	fmt.Printf(url, e.Exam, page)
	fmt.Printf("%s #%d\n", e.Exam, q)
}

func getPage(question int, pageSize int) int {
	if pageSize == 0 {
		pageSize = 5
	}
	return (question-1)/pageSize + 1
}
