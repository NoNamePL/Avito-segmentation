package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "postgreswd"
	dbname   = "Avito_segmentations"
)

func main() {
	psqlInfo := fmt.Sprintf("postgres://%s:%s@localhost:%d/%s?sslmode=disable", user, password, port, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := gin.Default()
	r.GET("/create", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "created",
		})
	})

	r.GET("/update",func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,gin.H{
			"message":"uploaded",
		})
	})

	r.GET("/delete", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,gin.H{
			"message":"deleted",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
