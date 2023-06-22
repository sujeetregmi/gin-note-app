package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/sujeetregmi/gin-note-app/controllers"
	"github.com/sujeetregmi/gin-note-app/middlewares"
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

	//session init
	store := memstore.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("notes", store))
	r.Use(middlewares.AuthenticateUser())

	r.GET("/notes", controllers.NotesIndex)
	r.GET("/notes/new", controllers.NotesNew)
	r.POST("/notes", controllers.NotesCreate)
	r.GET("/notes/:id", controllers.NotesShow)
	r.GET("/notes/edit/:id", controllers.NotesEditPage)
	r.POST("/notes/:id", controllers.NotesUpdate)
	r.DELETE("/notes/:id", controllers.NotesDelete)

	r.GET("/login", controllers.LoginPage)
	r.GET("/signup", controllers.SignupPage)

	r.POST("/login", controllers.Login)
	r.POST("/signup", controllers.Signup)
	r.POST("/logout", controllers.Logout)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/index.html", gin.H{
			"title":     "Notes Application",
			"logged_in": c.GetUint64("user_id") > 0,
		})
	})
	fmt.Println("Server Started")
	r.Run()
}
