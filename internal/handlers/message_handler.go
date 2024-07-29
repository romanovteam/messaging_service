package handlers

import (
	"github.com/gin-gonic/gin"
	"messaging_service/internal/models"
	"messaging_service/internal/services"
	"net/http"
)

type MessageHandler struct {
	service services.MessageService
}

func NewMessageHandler(service services.MessageService) *MessageHandler {
	return &MessageHandler{service: service}
}

func (h *MessageHandler) HandleMessage(c *gin.Context) {
	var message models.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON or missing 'content' field"})
		return
	}

	if err := h.service.ProcessMessage(&message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message processed successfully"})
}

func (h *MessageHandler) GetStatistics(c *gin.Context) {
	totalMessages, err := h.service.GetTotalMessagesCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	processedMessages, err := h.service.GetProcessedMessagesCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	unprocessedMessages, err := h.service.GetUnprocessedMessagesCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_messages":       totalMessages,
		"processed_messages":   processedMessages,
		"unprocessed_messages": unprocessedMessages,
	})
}
