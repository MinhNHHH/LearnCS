package exercises

import "io"

// LimitedReader wraps an io.Reader but stops after n bytes.
type LimitedReader struct {
	r io.Reader // underlying reader
	n int64     // max bytes remaining
}

// Read reads from the underlying reader but stops after n bytes.
func (lr *LimitedReader) Read(p []byte) (int, error) {
	if lr.n <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > lr.n {
		p = p[0:lr.n]
	}
	n, err := lr.r.Read(p)
	lr.n -= int64(n)
	return n, err
}

// LimitReader returns a Reader that reads from r but stops with EOF after n bytes.
func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{r, n}
}
