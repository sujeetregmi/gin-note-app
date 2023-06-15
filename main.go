package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sujeetregmi/gin-note-app/controllers"
	"github.com/sujeetregmi/gin-note-app/models"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Static("/vendor", "./static/vendor")
	r.LoadHTMLGlob("templates/**/**")

	// initialize db connection
	models.ConnectDatabase()
	models.DbMigrate()

	r.GET("/notes", controllers.NotesIndex)
	r.GET("/notes/new",controllers.NotesNew)
	r.POST("notes/",controllers.NotesCreate)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "views/index.html", gin.H{
			"title": "Notes App",
		})
	})
	r.Run()
}
