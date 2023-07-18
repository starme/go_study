/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"
	"os"
	"star/cmd"
)

func main() {
	mod := os.Getenv("APP_MODE")
	fmt.Sprintf("APP_MODE: %s", mod)

	cmd.Execute()
}

//func main() {
//	mux := asynq.NewServeMux()
//	router := gin.Default()
//	router.GET("/", func(c *gin.Context) {
//		time.Sleep(5 * time.Second)
//		c.String(http.StatusOK, "Welcome Gin Server")
//	})
//
//	srv := &http.Server{
//		Addr:    ":80",
//		Handler: router,
//	}
//	apiGroup := router.Group("/api")
//	{
//		apiGroup.GET("", func(c *gin.Context) {
//			queue.Dispatch(mux, hooks.ATask, hooks.ATaskOptions...)
//			c.JSON(200, gin.H{
//				"message": "api",
//			})
//		})
//	}
//
//	go func(mux *asynq.ServeMux) {
//		if err := queue.Start(mux); err != nil {
//			panic("job server is error")
//		}
//	}(mux)
//
//	go func() {
//		// 服务连接
//		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
//			log.Fatalf("listen: %s\n", err)
//		}
//	}()
//
//	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
//	quit := make(chan os.Signal)
//	signal.Notify(quit, os.Interrupt)
//	<-quit
//	log.Println("Shutdown Server ...")
//
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//	if err := srv.Shutdown(ctx); err != nil {
//		log.Fatal("Server Shutdown:", err)
//	}
//	log.Println("Server exiting")
//}
