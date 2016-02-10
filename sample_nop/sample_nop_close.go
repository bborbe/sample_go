package sample_nop_close

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var writer io.Writer = createWriter()
	var writerCloser io.WriteCloser = NopCloser(writer)
	needWriteCloser(writerCloser)
	fmt.Printf("done")
}

func createWriter() io.Writer {
	return bytes.NewBufferString("")
}

func needWriteCloser(writer io.WriteCloser) {
	writer.Close()
}

type nopCloser struct {
	io.Writer
}

func (nopCloser) Close() error { return nil }

func NopCloser(r io.Writer) io.WriteCloser {
	return nopCloser{r}
}
