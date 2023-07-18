package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"star/pkg/bus"
	"star/pkg/log"
)

type TListener struct{}

func (l TListener) Handler(event interface{}) {
	fmt.Printf("event: %v\n", event)
	log.Info("Handler: ", zap.Any("event", event))
}

func ApiV1(route *gin.RouterGroup) {
	handlers := func(c *gin.Context) {
		bus.Register("test", TListener{})
		bus.Dispatch("test", "this is test event")
		c.JSON(200, gin.H{
			"message": "api",
		})
	}
	//route.Use(func(c *gin.Context) {
	//	start := time.Now()
	//	path := c.Request.URL.Path
	//	query := c.Request.URL.RawQuery
	//	c.Next()
	//
	//	cost := time.Since(start)
	//	log.Channel("file").Debug(path,
	//		zap.Int("status", c.Writer.Status()),
	//		zap.String("method", c.Request.Method),
	//		zap.String("path", path),
	//		zap.String("query", query),
	//		zap.String("ip", c.ClientIP()),
	//		zap.String("user-agent", c.Request.UserAgent()),
	//		zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
	//		zap.Duration("cost", cost),
	//	)
	//})
	route.GET("", handlers)
}
