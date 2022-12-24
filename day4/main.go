package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id    int
	Title string
	Price float64
}

func main() {
	r := setupRouter()

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/book", bookHandler)

	r.Run(":8080")
}

func setupGin() *gin.Engine {
	return gin.Default()
}

func setupRouter() *gin.Engine {
	r := setupGin()

	setupGetBooksRouter(r)
	return r
}

func setupGetBooksRouter(r *gin.Engine) {
	r.GET("/books", func(c *gin.Context) {
		var books []Book
		books = append(books, Book{1, "Learn Go", 19.90})
		c.JSON(http.StatusOK, books)
	})
}

func setupDeleteBooksRouter(r *gin.Engine) {
	r.DELETE("/books/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, id)
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
	book := Book{1, "Learn Go", 19.90}

	result, err := json.Marshal(book)
	if err != nil {
		w.WriteHeader(500)
	}
	w.Write(result)
}
