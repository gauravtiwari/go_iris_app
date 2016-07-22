package main

import (
	"github.com/iris-contrib/template/django"
	"github.com/kataras/iris"
)

func main() {
	iris.UseTemplate(
		django.New(),
	).Directory("./templates", ".html")

	iris.Get("/", index)
	iris.Get("/hello", hello)
	iris.Listen(":3000")
}

func hello(ctx *iris.Context) {
	ctx.Render(
		"hello.html",
		map[string]interface{}{"Message": "Hello World"},
		iris.RenderOptions{"gzip": true},
	)
}

func index(ctx *iris.Context) {
	ctx.JSON(
		iris.StatusOK,
		map[string]string{"message": "Hello World"},
	)
}
