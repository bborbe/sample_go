package line_counter

type LineSplitter interface {
	SplitLines(content string) []string
}

type lineCounter struct{ lineSplitter LineSplitter }

func New(lineSplitter LineSplitter) *lineCounter {
	return &lineCounter{lineSplitter: lineSplitter}
}

func (l *lineCounter) CountLines(content string) int {
	var lines []string = l.lineSplitter.SplitLines(content)
	return len(lines)
}
