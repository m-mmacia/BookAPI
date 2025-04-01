package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

var books = []Book{
	{Id: "1", Title: "BookNumberOne", Author: "Bob", Price: 2.9},
	{Id: "2", Title: "BookNumberTwo", Author: "Bob", Price: 7.8},
	{Id: "3", Title: "BookNumberThree", Author: "Bob", Price: 10.23},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func addBook(c *gin.Context) {
	var newBook Book

	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Wrong data!"})
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func getBookId(c *gin.Context) {
	id := c.Param("id")
	for _, book := range books {
		if book.Id == id {
			c.IndentedJSON(http.StatusOK, book)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found :("})
}

func removeBookId(c *gin.Context) {
	id := c.Param("id")
	for i, book := range books {
		if book.Id == id {
			books = append(books[:i], books[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Book removed."})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found :("})
}

func main() {
	router := gin.Default()

	router.GET("/books", getBooks)
	router.POST("/books", addBook)
	router.GET("/books/:id", getBookId)
	router.DELETE("/books/:id", removeBookId)

	router.Run("localhost:5000") // change port
}
