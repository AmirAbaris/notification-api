package notification

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type NotificationHandler struct {
	notificationService *NotificationService
}

func NewNotificationHandler(s *NotificationService) *NotificationHandler {
	return &NotificationHandler{notificationService: s}
}

func (h *NotificationHandler) Create(c *gin.Context) {
	var body struct {
		UserID     uuid.UUID `json:"user_id"`
		TemplateID uuid.UUID `json:"template_id"`
		Status     string    `json:"status"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	createdNotification, err := h.notificationService.Create(
		c.Request.Context(),
		body.UserID,
		body.TemplateID,
		body.Status,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdNotification)
}

func (h *NotificationHandler) Get(c *gin.Context) {
	id := c.Param("id")

	notificationID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid notification id",
		})
		return
	}

	notification, err := h.notificationService.Get(
		c.Request.Context(),
		notificationID,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, notification)
}