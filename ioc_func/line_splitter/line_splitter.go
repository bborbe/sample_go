package line_splitter

import "strings"

type lineSplitter struct{}

func New() *lineSplitter { return new(lineSplitter) }

func (f *lineSplitter) SplitLines(content string) []string {
	return strings.Split(content, "\n")
}
