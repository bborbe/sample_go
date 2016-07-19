package main

import (
	"fmt"
	"net"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU() * 4)

	l, err := net.Listen("tcp", ":54321")
	defer l.Close()
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	counter := make(chan int, 10)
	go func() {
		i := 0
		for {
			i++
			counter <- i
		}
	}()

	for {
		rw, e := l.Accept()
		if e != nil {
			continue
		}
		go handleConn(rw, counter)
	}
}

func handleConn(c net.Conn, counter <-chan int) {
	fmt.Print(".")
	fmt.Fprintf(c, "HTTP/1.0 200 OK\n\n\ncount: %d\n", <-counter)
	c.Close()
}
