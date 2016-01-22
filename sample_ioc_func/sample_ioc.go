package main

import (
	"fmt"

	"github.com/bborbe/sample_go/sample_ioc_func/line_counter"
	"github.com/bborbe/sample_go/sample_ioc_func/line_splitter"
)

const CONTENT_TO_COUNT = `hello
world
`

func main() {
	lineSplitter := line_splitter.New()
	lineCounter := line_counter.New(lineSplitter.SplitLines)
	fmt.Printf("lines %d\n", lineCounter.CountLines(CONTENT_TO_COUNT))
}
