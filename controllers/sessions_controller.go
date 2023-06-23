package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sujeetregmi/gin-note-app/helpers"
	"github.com/sujeetregmi/gin-note-app/models"
)

func LoginPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"home/login.html",
		gin.H{},
	)
}

func SignupPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"home/signup.html",
		gin.H{},
	)
}

func Signup(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	confirmpassword := c.PostForm("confirm_password")

	// check if email already exists in database
	available := models.UserCheckAvailability(email)
	fmt.Println(available)
	if !available {
		c.HTML(
			http.StatusIMUsed,
			"home/signup.html",
			gin.H{
				"alert": "email already exists.",
			},
		)
		return
	}

	if password != confirmpassword {
		c.HTML(
			http.StatusNotAcceptable,
			"home/signup.html",
			gin.H{
				"alert": "Password MissMatched",
			},
		)
	}

	user := models.UserCreate(email, password)
	if user.ID == 0 {
		c.HTML(
			http.StatusNotAcceptable,
			"home/signup.html",
			gin.H{
				"alert": "Unable to create user!",
			},
		)
	} else {
		//signup successful , so set session
		helpers.SessionSet(c, user.ID)
		c.Redirect(
			http.StatusMovedPermanently,
			"/",
		)
	}
}

func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	user := models.UserCheck(email, password)
	if user != nil {
		//set sesion
		helpers.SessionSet(c, user.ID)
		c.Redirect(
			http.StatusMovedPermanently,
			"/",
		)
	} else {
		c.HTML(
			http.StatusOK,
			"home/login.html",
			gin.H{
				"alert": "Email or Password mis matched.!",
			},
		)
	}
}

func Logout(c *gin.Context) {
	// clears the session
	helpers.SessionClear(c)
	data := gin.H{
		"alert": "Logged Out",
	}
	c.HTML(
		http.StatusOK,
		"home/login.html",
		data,
	)
}
