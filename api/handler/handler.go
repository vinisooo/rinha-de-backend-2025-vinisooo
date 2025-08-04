package handler

import (
	"context"
	"net/http"
	"rinha_de_backend_vinisooo_2025/api/services"
	"rinha_de_backend_vinisooo_2025/config"

	"github.com/gin-gonic/gin"
)

func SetRoutes() *gin.Engine {
	redisClient := config.NewRedisClient()
	queueService := services.NewQueueService(redisClient)
	r := gin.Default()

	r.POST("/payments", func(c *gin.Context) {
		var requestData map[string]interface{}
		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		err := queueService.AddJob(context.Background(), "payments:processing", "process_payment", requestData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to queue payment", "details": err.Error()})
			return
		}

		c.JSON(http.StatusAccepted, gin.H{
			"message": "Payment queued for processing",
			"status":  "accepted",
		})
	})

	r.GET("/payments-summary", func(c *gin.Context) {
		// TODO: Implement
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, map[string]bool{
			"success": true,
		})
	})

	return r
}
