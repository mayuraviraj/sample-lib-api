package main

import "github.com/gin-gonic/gin"

type Book struct {
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
	ISBN   string  `json:"isbn"`
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/book", postBook)
	router.Run(":8080")
}
