package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"fhyx/Wechat/web"
)

func main() {
	var router *gin.Engine
	router = gin.Default()
	web.Strap(router)

	service := http.Server{
		Addr:         "127.0.0.1:8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	if err := service.ListenAndServe(); err != nil {
		log.Fatalf("error when start http service: %s", err)
	}
}
