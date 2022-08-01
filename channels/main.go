package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	ss1 := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	ss2, err := fast(ss1)
	if err != nil {
		log.Println(err)
	}
	for _, s := range ss2 {
		log.Printf("%q \n", s)
	}
}

func fast(ss1 []string) ([]string, error) {
	tasks := make(chan string)
	results := make(chan string)
	errors := make(chan error)

	// Send tasks to channel
	var wg1 sync.WaitGroup
	wg1.Add(1)
	go func() {
		defer wg1.Done()
		for _, s := range ss1 {
			tasks <- s
		}
	}()
	go func() {
		wg1.Wait()
		close(tasks)
	}()

	// Concurrent workers
	var wg2 sync.WaitGroup
	for i := 0; i < 3; i++ {
		i := i
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			cnt := 0
			for t := range tasks {
				results <- fmt.Sprintf("worker %v call %v - %s", i, cnt, t)
				time.Sleep(time.Millisecond)
				cnt++
				if cnt == 8 {
					errors <- fmt.Errorf("worker %v error", i)
				}
			}
		}()
	}
	go func() {
		wg2.Wait()
		close(results)
	}()

	// Receive results
	var ss2 []string
	var wg3 sync.WaitGroup
	wg3.Add(1)
	var err error
	go func() {
		defer wg3.Done()
		for {
			select {
			case e := <-errors:
				err = e
				return
			case s := <-results:
				ss2 = append(ss2, s)
			}
		}
	}()
	wg3.Wait()
	if err != nil {
		return nil, err
	}
	return ss2, nil
}
