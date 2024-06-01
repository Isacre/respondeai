package main

import (
	"net/http"
	"v1/internal/infra"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	wsAdapter := infra.NewWebSocketAdapter()
	wsService := infra.NewWebSocketService(wsAdapter)
	http.HandleFunc("/ws", wsService.HandleWebSocket)
	http.ListenAndServe(":8080", nil)

	r.Run()
}
