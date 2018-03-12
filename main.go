package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

// type MyType struct {
// }

// func main() {
// 	e := echo.New()
// 	e.Group("/admin")
// 	mt := MyType{}
// 	vof := reflect.ValueOf(mt)
// 	fmt.Println(vof.Type())
// }

func main() {
	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.GET("/", func(ctx echo.Context) error {
		sess, _ := session.Get("session", ctx)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400,
			HttpOnly: true,
		}
		// sess.Values["foo"] = "baz"
		// sess.Values["a"] = 101
		sess.Save(ctx.Request(), ctx.Response())
		fmt.Println(sess)
		return ctx.NoContent(http.StatusOK)
	})

	e.Logger.Fatal(e.Start(":3000"))
}
