package main

import (
	"fmt"
	"regexp"
)

func main() {
	r, _ := regexp.Compile("a (hello) b (world) c")

	result := r.FindAllStringSubmatch("a hello b world ca hello b world c", -1)
	fmt.Printf("%d\n", len(result))
	for _, r := range result {
		fmt.Printf("%v\n", r)
	}
}
