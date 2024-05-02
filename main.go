package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

}

func pushOrders(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "CI/CD successfully. test 5",
		"status":  "OK",
	})
}

func main() {
	r := gin.Default()

	// Thiết lập route
	r.GET("/push_orders", pushOrders)

	// Khởi chạy server
	r.Run(":5000")
}
