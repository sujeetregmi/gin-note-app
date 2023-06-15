package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sujeetregmi/gin-note-app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connectDatabase() {
	dsn := "host=localhost user=regmi password=regmi123 dbname=notes port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database.")
	}
	DB = database
}

func dbMigrate() {
	DB.AutoMigrate(&models.Note{})
}
func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Static("/vendor", "./static/vendor")
	r.LoadHTMLGlob("templates/**/**")

	// initialize db connection
	connectDatabase()
	dbMigrate()

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "views/index.html", gin.H{
			"title": "Notes App",
		})
	})
	r.Run()
}
