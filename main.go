package main

import (
	"fmt"
	"reflect"

	"gopkg.in/labstack/echo.v3"
)

type MyType struct {
}

func main() {
	e := echo.New()
	e.Group("/admin")
	mt := MyType{}
	vof := reflect.ValueOf(mt)
	fmt.Println(vof.Type())
}
