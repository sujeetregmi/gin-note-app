package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/sujeetregmi/gin-note-app/controllers"
	controllers_helper "github.com/sujeetregmi/gin-note-app/controllers/helpers"
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

	//grouping routes ->/notes
	notes := r.Group("/notes")
	{
		notes.GET("/", controllers.NotesIndex)
		notes.GET("/new", controllers.NotesNew)
		notes.POST("/", controllers.NotesCreate)
		notes.GET("/:id", controllers.NotesShow)
		notes.GET("/edit/:id", controllers.NotesEditPage)
		notes.POST("/:id", controllers.NotesUpdate)
		notes.DELETE("/:id", controllers.NotesDelete)
	}

	r.GET("/login", controllers.LoginPage)
	r.GET("/signup", controllers.SignupPage)

	r.POST("/login", controllers.Login)
	r.POST("/signup", controllers.Signup)
	r.POST("/logout", controllers.Logout)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/index.html", gin.H{
			"title":     "Notes Application",
			"logged_in": controllers_helper.IsUserLoggedIn(c),
		})
	})
	fmt.Println("Server Started")
	r.Run()
}
