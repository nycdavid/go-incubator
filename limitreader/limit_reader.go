package limitreader

import (
	"io"
)

type CappedReader struct {
	Limit int
	rdr   io.Reader
}

func (cpdRdr CappedReader) Read(p []byte) (int, error) {
	p = p[0:cpdRdr.Limit]
	read, err := cpdRdr.rdr.Read(p)
	if err != nil {
		return read, err
	}
	return cpdRdr.Limit, nil
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return CappedReader{Limit: int(n), rdr: r}
}
