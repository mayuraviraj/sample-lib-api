package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
}

type Book struct {
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
	ISBN   string  `json:"isbn"`
}

func main() {
	router := gin.Default()
	router.Use(LoggingMiddleware())
	router.GET("/books", getBooks)
	router.POST("/book", postBook)
	router.Run(":8080")
}

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("Incoming request",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("client_ip", c.ClientIP()),
			zap.String("user_agent", c.Request.UserAgent()),
		)

		c.Next() // Process request

		// Log response status after request processing
		logger.Info("Response sent",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
		)
	}
}
