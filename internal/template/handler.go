package template

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TemplateHandler struct {
	templateService TemplateService
}

func NewTemplateHandler(s TemplateService) *TemplateHandler {
	return &TemplateHandler{templateService: s}
}

func (h *TemplateHandler) Create(c *gin.Context) {
	var body struct {
		Name    string `json:"name"`
		Subject string `json:"subject"`
		Body    string `json:"body"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	createdUser, err := h.templateService.Create(
		c.Request.Context(),
		body.Name,
		body.Subject,
		body.Body,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdUser)

}

func (h *TemplateHandler) Get(c *gin.Context) {
	id := c.Param("id")

	uuid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user id",
		})
		return
	}

	user, err := h.templateService.Get(
		c.Request.Context(),
		uuid,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, user)

}
