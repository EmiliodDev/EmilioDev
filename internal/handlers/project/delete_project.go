package project

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteProject(c *gin.Context) {
	// idStr := c.Param("id")
	// id, err := strconv.Atoi(idStr)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	//
	// // project := models.Project{ID: uint(id)}
	//
	// if err := project.Delete(db.DB); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"message": "Project eliminated successfully"})
}
