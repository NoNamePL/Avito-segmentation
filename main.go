package main

import (
	"Avito_Intern/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)


func main() {

	 
	r := gin.Default()
	r.POST("/create", handlers.Create)

	r.POST("/update", handlers.Update)

	r.GET("/delete", handlers.Delete)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
