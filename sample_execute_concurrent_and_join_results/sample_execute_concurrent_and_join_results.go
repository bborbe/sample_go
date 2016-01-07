package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("print starts then all results are done")
	actions := []func() string{
		func() string {
			time.Sleep(2 * time.Second)
			return "action 2"
		},
		func() string {
			time.Sleep(3 * time.Second)
			return "action 3"
		},
		func() string {
			time.Sleep(4 * time.Second)
			return "action 4"
		},
		func() string {
			time.Sleep(1 * time.Second)
			return "action 1"
		},
	}
	for _, message := range executeConcurrent(actions) {
		fmt.Printf("%s\n", message)
	}
	fmt.Println("done")
}

func executeConcurrent(actions []func() string) []string {
	var wg sync.WaitGroup
	var list []string
	results := make(chan string)
	done := make(chan bool)
	go func() {
		for result := range results {
			list = append(list, result)
		}
		done <- true
	}()
	for _, a := range actions {
		wg.Add(1)
		// copy action
		action := a
		go func() {
			results <- action()
			wg.Done()
		}()
	}
	wg.Wait()
	close(results)

	<-done

	return list
}
