package limitreader

import (
	"io"
)

type CappedReader struct {
	Limit int
}

func (cpdRdr CappedReader) Read(p []byte) (int, error) {
	return 0, nil
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return CappedReader{Limit: int(n)}
}
