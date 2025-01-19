package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func getBooks(c *gin.Context) {
	logger.Info("fetching books")
	c.IndentedJSON(http.StatusOK, sampleBooks)
}

func postBook(c *gin.Context) {
	var newBook Book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	sampleBooks = append(sampleBooks, newBook)
	logger.Info("New book added", zap.String("title", newBook.Title))
	c.IndentedJSON(http.StatusCreated, newBook)
}
