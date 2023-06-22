package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sujeetregmi/gin-note-app/controllers/helpers"
	"github.com/sujeetregmi/gin-note-app/models"
)

func NotesIndex(c *gin.Context) {
	currentUser := helpers.GetUserFromRequest(c)
	if currentUser == nil || currentUser.ID == 0 {
		c.HTML(
			http.StatusUnauthorized,
			"notes/index.html",
			gin.H{
				"alert": "Unautorized Access!",
			},
		)
		return
	}
	notes := models.NotesAll(currentUser)
	c.HTML(
		http.StatusOK,
		"notes/index.html",
		gin.H{
			"notes": notes,
		},
	)
}

func NotesNew(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"notes/new.html",
		gin.H{},
	)
}

func NotesCreate(c *gin.Context) {
	currentUser := helpers.GetUserFromRequest(c)
	if currentUser == nil || currentUser.ID == 0 {
		c.HTML(
			http.StatusUnauthorized,
			"notes/index.html",
			gin.H{
				"alert": "Unautorized Access!",
			},
		)
		return
	}
	name := c.PostForm("name")
	content := c.PostForm("content")
	models.NoteCreate(currentUser, name, content)
	c.Redirect(http.StatusMovedPermanently, "/notes")
}

func NotesShow(c *gin.Context) {
	currentUser := helpers.GetUserFromRequest(c)
	if currentUser == nil || currentUser.ID == 0 {
		c.HTML(
			http.StatusUnauthorized,
			"notes/index.html",
			gin.H{
				"alert": "Unautorized Access!",
			},
		)
		return
	}
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error:%v", err)
	}
	note := models.NotesFind(currentUser, id)
	c.HTML(
		http.StatusOK,
		"notes/show.html",
		gin.H{
			"note": note,
		},
	)
}

func NotesEditPage(c *gin.Context) {
	currentUser := helpers.GetUserFromRequest(c)
	if currentUser == nil || currentUser.ID == 0 {
		c.HTML(
			http.StatusUnauthorized,
			"notes/index.html",
			gin.H{
				"alert": "Unautorized Access!",
			},
		)
		return
	}
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error:%v", err)
	}
	note := models.NotesFind(currentUser, id)
	c.HTML(
		http.StatusOK,
		"notes/edit.html",
		gin.H{
			"note": note,
		},
	)
}

func NotesUpdate(c *gin.Context) {
	currentUser := helpers.GetUserFromRequest(c)
	if currentUser == nil || currentUser.ID == 0 {
		c.HTML(
			http.StatusUnauthorized,
			"notes/index.html",
			gin.H{
				"alert": "Unautorized Access!",
			},
		)
		return
	}
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error:%v", err)
	}
	note := models.NotesFind(currentUser, id)
	name := c.PostForm("name")
	content := c.PostForm("content")
	note.Update(name, content)
	c.Redirect(http.StatusMovedPermanently, "/notes/"+idStr)
}

func NotesDelete(c *gin.Context) {
	currentUser := helpers.GetUserFromRequest(c)
	if currentUser == nil || currentUser.ID == 0 {
		c.HTML(
			http.StatusUnauthorized,
			"notes/index.html",
			gin.H{
				"alert": "Unautorized Access!",
			},
		)
		return
	}
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error:%v", err)
	}
	models.NotesMarkDelete(currentUser, id)
	c.Redirect(http.StatusSeeOther, "/notes")
}
