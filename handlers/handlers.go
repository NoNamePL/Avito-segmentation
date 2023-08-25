package handlers

import (
	"Avito_Intern/db"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSONError(c *gin.Context, err error, bcode string, code int) {
	
    c.Header("Content-Type", "application/json; charset=utf-8")
    c.Header("X-Content-Type-Options", "nosniff")
    c.Writer.WriteHeader(code)
    json.NewEncoder(c).Encode(struct{
        Code    string
        Message string
    }{
        Code:    bcode,
        Message: err.Error(),
    })
}

func Create(c *gin.Context) {
	db, err := db.Init()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	decoder := json.NewDecoder(c.Request.Body)

	var segmentations db.Segmentations

	err := decoder.Decode(&segmentations)
	if err != nil {
		JSONError(c,fmt.Errorf("cannot unmarshal segmentations: %w", err),"custom_code",http.StatusUnprocessableEntity)
		return 
	}
	for idx := range segmentations{
		
	}
	// rows, err := db.Query("INSERT INTO segmentations VALUE (DEFAULT,'{%s}')")

	c.JSON(http.StatusOK, gin.H{
		"message": "created",
	})
}
func Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "updated",
	})
}

func Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "deleted",
	})
}
