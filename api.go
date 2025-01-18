package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getBooks(c *gin.Context) {
	fmt.Println("getBooks")
	fmt.Println(sampleBooks)
	c.IndentedJSON(http.StatusOK, sampleBooks)
}

func postBook(c *gin.Context) {
	var newBook Book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	sampleBooks = append(sampleBooks, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}
