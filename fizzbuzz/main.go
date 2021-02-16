package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("need one parameter")
		os.Exit(1)
	}
	param := os.Args[1]
	number, err := strconv.Atoi(param)
	if err != nil {
		fmt.Printf("parse parameter failed: %v", err)
		os.Exit(1)
	}
	if number%3 == 0 && number%5 == 0{
		fmt.Println("FizzBuzz")
		return
	}
	if number%3 == 0 {
		fmt.Println("Fizz")
		return
	}
	if number%5 == 0 {
		fmt.Println("Buzz")
		return
	}
	fmt.Println(param)
}
