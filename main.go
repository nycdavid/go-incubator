package main

import (
	"fmt"
	"reflect"

	"gopkg.in/labstack/echo.v3"
)

type MyType struct {
	Id   string
	Name string
}

func (mt *MyType) Foob() {
	fmt.Println("Called foob")
}

func main() {
	e := echo.New()
	e.Group("/admin")
	mt := &MyType{}
	vof := reflect.ValueOf(mt)
	vof.MethodByName("Foob").Call([]reflect.Value{})
}
