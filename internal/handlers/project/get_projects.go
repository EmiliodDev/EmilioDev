package project

import (
	"net/http"

	"github.com/EmiliodDev/EmilioDev/internal/db"
	"github.com/EmiliodDev/EmilioDev/internal/models"
	"github.com/gin-gonic/gin"
)

func GetProjects(c *gin.Context) {
	projects, err := models.GetAllProjects(db.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, projects)
}
