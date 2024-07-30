package router

import (
	"github.com/EmiliodDev/EmilioDev/internal/handlers"
	"github.com/EmiliodDev/EmilioDev/internal/handlers/project"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("internal/templates/*")

	r.Static("/static", "./static")

	r.GET("/", handlers.Home)

	api := r.Group("/projects")
	{
		api.POST("", project.CreateProject)
		api.PUT("/:id", project.UpdateProject)
		api.DELETE("/:id", project.DeleteProject)
		api.GET("", project.GetProjects)
	}

	return r
}
