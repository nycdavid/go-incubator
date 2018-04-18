package writer

import (
	"testing"
)

type ClassicRW struct {
}

func (crw ClassicRW) Read(p []byte) (int, error) {
	return 0, nil
}

func (crw ClassicRW) Write(p []byte) (int, error) {
	return 0, nil
}

func TestRWFuncAcceptsReadWriter(t *testing.T) {
	var myRW ClassicRW
	RW(myRW)
}
