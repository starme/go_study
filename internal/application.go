package internal

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"star/pkg/config"
	"time"
)

type Application struct {
	webserver *http.Server

	handle *gin.Engine
}

func (app *Application) boot() {
	gin.DisableConsoleColor()
	app.handle = gin.Default()
	app.webserver = &http.Server{
		Addr:    fmt.Sprintf(":%d", config.GetInt("APP_PORT")),
		Handler: app.handle,
	}

}

func (app *Application) Run(providers ...func(*Application)) {
	app.boot()

	for _, provider := range providers {
		provider(app)
		//log.Printf("provider %d", i)
	}

	go func() {
		// 服务连接
		err := app.webserver.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.webserver.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

func (app *Application) GetRoute() *gin.Engine {
	return app.handle
}
