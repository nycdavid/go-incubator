package suitetesting

import (
	"fmt"
)

type Thinger interface {
	DoThing()
}

type OneThinger struct{}

func (t *OneThinger) DoThing() {
	fmt.Println("One Thinger")
}

type TwoThinger struct{}

func (t *TwoThinger) DoThing() {
	fmt.Println("One Thinger")
}
