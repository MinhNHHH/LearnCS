package exercises

import "io"

type CountingWriters struct {
	writer io.Writer
	count  int64
}

func (cw *CountingWriters) Write(p []byte) (int, error) {
	n, err := cw.writer.Write(p)
	cw.count += int64(n)
	return n, err
}

func (c *CountingWriters) CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := &CountingWriters{writer: w}
	return cw, &cw.count
}
