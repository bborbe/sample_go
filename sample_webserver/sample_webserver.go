package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"
)

type helloHandler struct {
	counter <-chan int
}

func NewHelloHandler(counter <-chan int) *helloHandler {
	h := new(helloHandler)
	h.counter = counter
	return h
}

func (h *helloHandler) ServeHTTP(res http.ResponseWriter, request *http.Request) {
	fmt.Print(".")
	fmt.Fprintf(res, "count: %d\n", <-h.counter)
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU() * 4)

	counter := make(chan int, 10)
	go func() {
		i := 0
		for {
			i++
			counter <- i
		}
	}()

	log.Printf("%s webserver started", time.Now().Format("2006-01-02T15:04:05"))
	handler := NewHelloHandler(counter)
	err := http.ListenAndServe(":54321", handler)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return
	}
}
