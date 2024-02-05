package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	x := getValueInsideParentheses("projects/test/topics/blabla")
	fmt.Println(x)
}

func getValueInsideParentheses(s string) string {
	if !strings.Contains(s, "projects/my-project-") {
		x := regexExp.FindStringSubmatch(s)
		return x[1]
	}
	return "no match"
}

var regexExp *regexp.Regexp

func init() {
	regexExp = regexp.MustCompile(`^projects/(.+)/topics/.+$`)
}
