package suitetesting

import (
	"testing"
)

func Suite(t *testing.T, impl Thinger) {
	impl.DoThing()
}

func TestThinger_thingerOne(t *testing.T) {
	o := &OneThinger{}
	Suite(t, o)
}

func TestThinger_thingerTwo(t *testing.T) {
	o := &TwoThinger{}
	Suite(t, o)
}
