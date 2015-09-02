package controllers

import (
  "dmas/models"
	"strconv"

  "github.com/gin-gonic/gin"
  "gopkg.in/mgo.v2/bson"
)

func GetAllQuestions(c *gin.Context) {
	result := models.GetAllQuestions()

	c.JSON(200, result)
}

func GetOneQuestion(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 0, 64)

	if err != nil {
		c.JSON(500, gin.H{"There was some problem...": err})
	}

	result := models.GetQuestionByID(id)
	c.JSON(200, result)
}

func CreateQuestion(c *gin.Context) {
	title := c.PostForm("title")
	answer := c.PostForm("answer")
	a1 := c.PostForm("a1")
	a2 := c.PostForm("a2")
	a3 := c.PostForm("a3")
	a4 := c.PostForm("a4")

	status, err := models.CreateQuestion(title, answer, a1, a2, a3, a4)

	if status {
		c.JSON(200, gin.H{"status": "All good"})
	} else {
		c.JSON(500, gin.H{"There was some problem...": err})
	}
}

func UpdateQuestion(c *gin.Context) {
	id, err := strconv.ParseInt(c.PostForm("id"), 0, 64)

	if err != nil {
		c.JSON(500, gin.H{"There was some problem with id...": err})
	}

	updateParams := bson.M{}

	if c.PostForm("title") != "" {
		updateParams["title"] = c.PostForm("title")
	}

	if c.PostForm("answer") != "" {
		updateParams["answer"] = c.PostForm("answer")
	}

	if c.PostForm("a1") != "" {
		updateParams["a1"] = c.PostForm("a1")
	}

	if c.PostForm("a2") != "" {
		updateParams["a2"] = c.PostForm("a2")
	}

	if c.PostForm("a3") != "" {
		updateParams["a3"] = c.PostForm("a3")
	}

	if c.PostForm("a4") != "" {
		updateParams["a4"] = c.PostForm("a4")
	}

	status, err := models.UpdateQuestion(id, updateParams)

	if status {
		c.JSON(200, gin.H{"status": "All good"})
	} else {
		c.JSON(500, gin.H{"There was some problem...": err})
	}
}

func DeleteQuestion(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 0, 64)

	if err != nil {
		c.JSON(500, gin.H{"There was some problem...": err})
	}

	deleted := models.DeleteQuestion(id)
	if deleted {
		c.JSON(200, gin.H{"status": "Document is deleted"})
	} else {
		c.JSON(500, gin.H{"status": "There was some problem"})
	}

}
