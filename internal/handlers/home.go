package handlers

import (
	"net/http"

	"github.com/EmiliodDev/EmilioDev/internal/templates"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {

	c.Header("Content-Type", "text/html")
	err := templates.Home().Render(c, c.Writer)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error rendering the template")
	}
}
