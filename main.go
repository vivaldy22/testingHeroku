package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	g.GET("/", func(ctx *gin.Context) {
		ctx.Writer.Write([]byte("Hai"))
	})

	g.Run()
}
