package handlers

import (
	"Avito_Intern/db"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Обработка ошибки JSON
func JSONError(c *gin.Context, err error, bcode string, code int) {

	c.Header("Content-Type", "application/json; charset=utf-8")
	c.Header("X-Content-Type-Options", "nosniff")
	c.Writer.WriteHeader(code)
	json.NewEncoder(c.Writer).Encode(struct {
		Code    string
		Message string
	}{
		Code:    bcode,
		Message: err.Error(),
	})
}

// Реализация JSON запроса
type Segment struct {
	userID int      `json:"user_id"`
	item   []string `json:"item"`
}

type Segmentations []Segment

// Запрос к БД на создание сегмента
func Create(c *gin.Context) {
	db, err := db.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	decoder := json.NewDecoder(c.Request.Body)

	var segmentations Segment // Segmentations

	err = decoder.Decode(&segmentations)
	if err != nil {
		JSONError(c, fmt.Errorf("cannot unmarshal segmentations: %w", err), "custom_code", http.StatusUnprocessableEntity)
		return
	}
	// log.Println(segmentations.item)
	// добавление в БД
	// for _, val := range segmentations { //segmentations.Segments
	// обработка запроса
	item := strings.ReplaceAll(segmentations.item[0], ",", "\", \"")
	// log.Println("privet")
	query := fmt.Sprintf("INSERT INTO segmentations VALUE (DEFAULT,'{%s}')", item) // val.Item
	rows, err := db.PrepareContext(c, query)
	// rows, err := db.Query("INSERT INTO segmentations VALUE (DEFAULT,'{%s}')", val.Item)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	// }
	// Ответ при успешном создании
	c.JSON(http.StatusOK, gin.H{
		"message": "created",
	})
}

// Обновление данных
func Update(c *gin.Context) {

	db, err := db.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	decoder := json.NewDecoder(c.Request.Body)

	var segmentations Segmentations

	err = decoder.Decode(&segmentations)
	if err != nil {
		JSONError(c, fmt.Errorf("cannot unmarshal segmentations: %w", err), "custom_code", http.StatusUnprocessableEntity)
		return
	}

	for _, val := range segmentations {
		query := fmt.Sprintf("SELECT * FROM segmentations WHERE user_id = %d", val.userID)
		rows, err := db.QueryContext(c, query)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			var (
				user_id int
				item    string
			)
			if err := rows.Scan(&user_id, &item); err != nil {
				log.Fatal(err)
			}
			query := fmt.Sprintf("UPDATE segmentations SET segment = '{%s}' WHERE user = %d", item, user_id)
			rows, err := db.QueryContext(c, query)
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()
		}
	}
	// Ответ при успешном изменении
	c.JSON(http.StatusOK, gin.H{
		"message": "updated",
	})
}

// Удаление пользователя
func Delete(c *gin.Context) {

	db, err := db.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	decoder := json.NewDecoder(c.Request.Body)

	var segmentations Segmentations

	err = decoder.Decode(&segmentations)
	if err != nil {
		JSONError(c, fmt.Errorf("cannot unmarshal segmentations: %w", err), "custom_code", http.StatusUnprocessableEntity)
		return
	}

	for _, val := range segmentations {
		query := fmt.Sprintf("DELETE FROM segmentations WHERE user_id = %d", val.userID)
		rows, err := db.QueryContext(c, query)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
	}
	// Ответ при успешном удалении
	c.JSON(http.StatusOK, gin.H{
		"message": "deleted",
	})
}

// удаление сегмента
func DeleteSegment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "deleted Segment",
	})
}
