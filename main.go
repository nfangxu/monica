package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()
	app.Use(iris.Compression)

	app.Get("/", func(ctx iris.Context) {
		ctx.JSON(map[string]any{
			"code": "0x20240128",
			"data": map[string]any{
				"slogon": "Hello, Monica",
			},
		})
	})

	app.Listen(":8080")
}
