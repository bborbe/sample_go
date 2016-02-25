package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("print starts immediately")
	var wg sync.WaitGroup
	wg.Add(1)
	go action("action 2", 2, &wg)
	wg.Add(1)
	go action("action 4", 4, &wg)
	wg.Add(1)
	go action("action 3", 3, &wg)
	wg.Add(1)
	go action("action 1", 1, &wg)
	wg.Wait()
	fmt.Println("done")
}

func action(message string, timeToSleep time.Duration, wg *sync.WaitGroup) {
	time.Sleep(timeToSleep * time.Second)
	fmt.Println(message)
	wg.Done()
}
