package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *UserService
}

func NewUserHandler(s *UserService) *UserHandler {
	return &UserHandler{userService: s}
}

func (h *UserHandler) Create(c *gin.Context) {
	var body struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	createdUser, err := h.userService.Create(
		c.Request.Context(),
		body.Name,
		body.Email,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdUser)

}
