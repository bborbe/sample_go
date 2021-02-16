package main

import (
	"flag"
	"fmt"
	"runtime"
)

const bufferSize = 1

func main() {
	valuePtr := flag.Int64("n", 20, "number of primes to generate")
	concurrentPtr := flag.Int("c", runtime.NumCPU(), "number of go routines")
	flag.Parse()
	concurrent := *concurrentPtr
	runtime.GOMAXPROCS(*concurrentPtr)
	value := *valuePtr
	counterChan := make(chan int64, bufferSize)
	primeChan := make(chan int64, bufferSize)

	go genCounter(counterChan)
	for i := 0; i < concurrent; i++ {
		go genPrimes(counterChan, primeChan)
	}

	for i := value; i > 0; i-- {
		fmt.Println(<-primeChan)
	}
}

func genCounter(counterChan chan<- int64) {
	i := int64(1)
	for {
		i++
		counterChan <- i
	}
}

func genPrimes(counterChan <-chan int64, primeChan chan<- int64) {
	for {
		i := <-counterChan
		if isPrime(i) {
			primeChan <- i
		}
	}
}

func isPrime(n int64) bool {
	for i := int64(2); i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
