package routes

import (
	"github.com/gin-gonic/gin"
)

func ApiV1(route *gin.RouterGroup) {
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
	route.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "api",
		})
	})
}
