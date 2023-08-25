package main

import (
	"Avito_Intern/handlers"

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

	// db, err := db.Init()
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()

	r := gin.Default()
	r.POST("/create", handlers.Create)

	r.POST("/update", handlers.Update)

	r.GET("/delete", handlers.Delete)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
