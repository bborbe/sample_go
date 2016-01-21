package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("print starts immediately")
	maxConcurrentGoRoutines := 4
	throttle := make(chan bool, maxConcurrentGoRoutines)
	var wg sync.WaitGroup
	for i := 0; i < 40; i++ {
		message := fmt.Sprintf("action %d", i)
		wg.Add(1)
		go action(message, &wg, throttle)
	}
	wg.Wait()
	fmt.Println("done")
}

func action(message string, wg *sync.WaitGroup, throttle chan bool) {
	throttle <- true
	time.Sleep(1000 * time.Millisecond)
	fmt.Printf("%s %s\n", time.Now().Format("15:04:05"), message)
	wg.Done()
	<-throttle
}
