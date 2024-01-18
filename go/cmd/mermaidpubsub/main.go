package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/iterator"
)

func main() {
	lines, err := listTopics("my-project")
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create("pubsub.md")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.WriteString("```mermaid\n")
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.WriteString("flowchart TB\n")
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range lines {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	_, err = f.WriteString("```\n")
	if err != nil {
		log.Fatal(err)
	}
}

func listTopics(projectID string) ([]string, error) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("pubsub.NewClient: %w", err)
	}
	defer client.Close()
	var topics []*pubsub.Topic
	it := client.Topics(ctx)
	for {
		topic, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("Next: %w", err)
		}
		topics = append(topics, topic)
	}
	var lines []string
	for _, t := range topics {
		t := t
		topicName := simplify(t.String())
		it := t.Subscriptions(ctx)
		for {
			s, err := it.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return nil, fmt.Errorf("Next: %w", err)
			}
			subName := simplify(s.String())
			newL := fmt.Sprintf("%s-->%s;", topicName, subName)
			existAlready := false
			for _, l := range lines {
				l := l
				if l == newL {
					existAlready = true
					break
				}
			}
			if len(lines) < 1 || !existAlready {
				lines = append(lines, newL)
			}
			fmt.Println(newL)
		}
	}
	return lines, nil
}

func simplify(s string) string {
	return overwrite(replaceSubstrings((untilFirstUnderscore(afterLastSlash(s)))))
}

func afterLastSlash(s string) string {
	s2 := strings.Split(s, "/")
	return s2[len(s2)-1]
}

func untilFirstUnderscore(s string) string {
	s2 := strings.Split(s, "_")
	return s2[0]
}

func replaceSubstrings(s string) string {
	sOut := s
	subs := []string{
		"-dev",
		"-topic",
	}
	for _, sub := range subs {
		sub := sub
		sOut = strings.ReplaceAll(sOut, sub, "")
	}
	return sOut
}

func overwrite(s string) string {
	if strings.Contains(s, "my-service") {
		return "my-service"
	}
	return s
}
