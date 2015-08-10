package controllers

import (
  "dmas/models"

  "github.com/gin-gonic/gin"
  "github.com/asaskevich/govalidator"
  "gopkg.in/mgo.v2/bson"
)

//LoginHandler handle simple route for login - return template or something
func LoginHandler(c *gin.Context) {
	c.String(200, "Login route here")
}

//RegisterHandler handle users registrations
func RegisterHandler(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")

	//we are using external library to validate email and other things if needed
	if govalidator.IsEmail(email) {
		//c.String(200, "Username is: %s and e-mail %s", username, email)
		//or we can return JSON as well
		c.JSON(200, gin.H{"username": username, "email": email})
	} else {
		//email is not valid we can't register this user
		c.String(400, "We are sorry but email: %s is not valid e-mail adress", email)
	}
}

// GetUsers returns all users from database
func GetAllUsers(c *gin.Context) {
	result := models.GetAllUsers()
	c.JSON(200, result)
}

// GetOneUser returns single users details
func GetOneUser(c *gin.Context) {
	id := c.Param("id")

	result := models.GetUserById(id)
	c.JSON(200, result)
}

// CreateUser is inserting new user
func CreateUser(c *gin.Context) {

	var data = map[string]string{}
	data["firstname"] = c.PostForm("firstname")
	data["lastname"] = c.PostForm("lastname")
	data["username"] = c.PostForm("username")
	data["email"] = c.PostForm("email")
	data["facebook_id"] = c.PostForm("facebook_id")
	data["is_active"] = c.PostForm("active")

	status, err := models.CreateUser(data)

	if status {
		c.JSON(200, gin.H{"status": "All good"})
	} else {
		c.JSON(500, gin.H{"There was some problem...": err})
	}
}

// UpdateUser updates single user
func UpdateUser(c *gin.Context) {
	id := c.PostForm("id")
	updateParams := bson.M{}

	if c.PostForm("firstname") != "" {
		updateParams["firstname"] = c.PostForm("firstname")
	}

	if c.PostForm("lastname") != "" {
		updateParams["lastname"] = c.PostForm("lastname")
	}

	if c.PostForm("username") != "" {
		updateParams["username"] = c.PostForm("username")
	}

	if c.PostForm("email") != "" {
		updateParams["email"] = c.PostForm("email")
	}

	if c.PostForm("facebook_id") != "" {
		updateParams["facebook_id"] = c.PostForm("facebook_id")
	}

	if c.PostForm("active") != "" {
		updateParams["active"] = c.PostForm("active")
	}

	status, err := models.UpdateUser(id, updateParams)

	if status {
		c.JSON(200, gin.H{"status": "All good"})
	} else {
		c.JSON(500, gin.H{"There was some problem...": err})
	}
}

// DeleteUser removes user
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	deleted := models.DeleteUser(id)
	if deleted {
		c.JSON(200, gin.H{"status": "User is deleted"})
	} else {
		c.JSON(500, gin.H{"status": "There was some problem"})
	}

}
