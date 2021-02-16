package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool, 1)
	fmt.Println("start action")
	go action(done)
	fmt.Println("wait on action done")
	<-done
	fmt.Println("finished")
}

func action(done chan<- bool) {
	time.Sleep(5 * time.Second)
	fmt.Println("done")
	done <- true
}
