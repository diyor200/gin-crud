package main

import (
	"github.com/diyor200/golang-gin-poc/books"
	"github.com/diyor200/golang-gin-poc/common/db"
	"github.com/gin-gonic/gin"
)

func main() {
	const (
		port  = "8000"
		dburl = "postgres://postgres:2001@localhost:5432/postgres"
	)

	router := gin.Default()
	dbHandler := db.Init(dburl)

	books.RegisterRoutes(router, dbHandler)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"port":  port,
			"dburl": dburl,
		})
	})

	router.Run("localhost:" + port)
}
