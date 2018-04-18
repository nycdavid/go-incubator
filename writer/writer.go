package writer

import (
	"io"
)

func RW(rw io.ReadWriter) {
	foo := []byte{}
	rw.Read(foo)
	rw.Write(foo)
}
