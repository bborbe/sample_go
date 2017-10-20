package main

import (
	"fmt"
)

func main() {
	a := "hello world"
	fmt.Printf("%v\n", a[0: 0])
	fmt.Printf("%v\n", a[0: 1])
	fmt.Printf("%v\n", a[0: 2])
	fmt.Printf("%v\n", a[1: 1])
	fmt.Printf("%v\n", a[3: 4])
	fmt.Printf("%v\n", a[0: len(a)])
}
