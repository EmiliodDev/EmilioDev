package handlers

import (
	"net/http"

	"github.com/EmiliodDev/EmilioDev/internal/db"
	"github.com/EmiliodDev/EmilioDev/internal/models"
	"github.com/gin-gonic/gin"
)

func Projects(c *gin.Context) {
	projects, err := models.GetAllProjects(db.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "projects.templ", gin.H{
		"Title":    "Projects",
		"Projects": projects,
	})
}
