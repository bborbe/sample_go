package line_counter

type SplitLines func(string) []string

type lineCounter struct{ splitLines SplitLines }

func New(splitLines SplitLines) *lineCounter {
	return &lineCounter{splitLines: splitLines}
}

func (l *lineCounter) CountLines(content string) int {
	var lines []string = l.splitLines(content)
	return len(lines)
}
