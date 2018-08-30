package main

import (
	"github.com/asmexie/go-logger/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	logger.SetConsole(true)
	logger.SetRollingDaily("./logs", "server.log")
	logger.Info("Server Listen at port:8080")
	router := gin.Default()
	router.GET("/api/ipapp/taobao", iptaobao)
	router.Run(":8080")
}
