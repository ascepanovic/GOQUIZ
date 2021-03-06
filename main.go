package main

import (
	"dmas/models"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	port := ":4555" //on this port our web server will be liste for incoming requests

	r := gin.Default() //default instance of gin framework used as router as well
	//CORSMiddlware will inject needed headers for our angular client
	r.Use(dmas.CORSMiddleware())

	//for sockets
	r.GET("/ws", func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	})

	//listen for certain paths - Home
	r.GET("/", HomeHandler)

	//now we will do some routes grouping and there we are going to have basic post/get functions/handlers
	userRoutes := r.Group("/user")
	{
		userRoutes.POST("/register", registerHandler)
		userRoutes.GET("/login", loginHandler)
	}

	//and testing some api routes now
	apiRoutes := r.Group("api")
	{
		apiRoutes.GET("/fetch/:link", apiFetch)
		apiRoutes.GET("/json", apiTournaments)

		// API routes for users
		apiRoutes.GET("/users", apiUsers)
		apiRoutes.GET("/user/:id", apiOneUser)
		apiRoutes.POST("/user/create", apiCreateUser)
		apiRoutes.POST("/user/update", apiUpdateUser)
		apiRoutes.POST("/user/delete/:id", apiDeleteUser)

		// API routes for questions
		apiRoutes.GET("/questions", apiQuestions)
		apiRoutes.GET("/question/:id", apiOneQuestion)
		apiRoutes.POST("/question/create", apiCreateQuestion)
		apiRoutes.POST("/question/update", apiUpdateQuestion)
		apiRoutes.POST("/question/delete/:id", apiDeleteQuestion)
	}

	//now listen on defined port
	r.Run(port)
}

//socket structure
var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

//socket handler
func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(t, msg)
	}
}

//HomeHandler this function is using ginContext
func HomeHandler(c *gin.Context) {
	//and sand some response do
	c.String(200, "Hello from gin")
}

//handle simple route for login - return template or something
func loginHandler(c *gin.Context) {
	c.String(200, "Login route here")
}

//for now just return submited params
func registerHandler(c *gin.Context) {
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

//let's for learning purposes see how to interact for example with rest api
func apiTournaments(c *gin.Context) {
	response, err := http.Get("http://jsonplaceholder.typicode.com/posts")

	if err != nil {
		c.JSON(400, gin.H{"error": err})
	} else {
		contents, e := ioutil.ReadAll(response.Body) //read complete body if any with utils

		if e != nil {
			c.JSON(400, gin.H{"We have an error": e})
		} else {
			c.String(200, string(contents))
		}
	}
}

func apiUsers(c *gin.Context) {
	result := dmas.GetAllUsers()
	c.JSON(200, result)
}

func apiOneUser(c *gin.Context) {
	id := c.Param("id")

	result := dmas.GetUserById(id)
	c.JSON(200, result)
}

func apiCreateUser(c *gin.Context) {

	var data = map[string]string{}
	data["firstname"] = c.PostForm("firstname")
	data["lastname"] = c.PostForm("lastname")
	data["username"] = c.PostForm("username")
	data["email"] = c.PostForm("email")
	data["facebook_id"] = c.PostForm("facebook_id")
	data["is_active"] = c.PostForm("active")

	status, err := dmas.CreateUser(data)

	if status {
		c.JSON(200, gin.H{"status": "All good"})
	} else {
		c.JSON(500, gin.H{"There was some problem...": err})
	}
}

func apiUpdateUser(c *gin.Context) {
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

	status, err := dmas.UpdateUser(id, updateParams)

	if status {
		c.JSON(200, gin.H{"status": "All good"})
	} else {
		c.JSON(500, gin.H{"There was some problem...": err})
	}
}

func apiQuestions(c *gin.Context) {
	fmt.Println("Into questions function...")
	result := dmas.GetAllQuestions()

	c.JSON(200, result)
}

func apiOneQuestion(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 0, 64)

	if err != nil {
		c.JSON(500, gin.H{"There was some problem...": err})
	}

	result := dmas.GetQuestionByID(id)
	c.JSON(200, result)
}

func apiDeleteUser(c *gin.Context) {
	id := c.Param("id")

	deleted := dmas.DeleteUser(id)
	if deleted {
		c.JSON(200, gin.H{"status": "User is deleted"})
	} else {
		c.JSON(500, gin.H{"status": "There was some problem"})
	}

}

func apiCreateQuestion(c *gin.Context) {
	title := c.PostForm("title")
	answer := c.PostForm("answer")
	a1 := c.PostForm("a1")
	a2 := c.PostForm("a2")
	a3 := c.PostForm("a3")
	a4 := c.PostForm("a4")

	status, err := dmas.CreateQuestion(title, answer, a1, a2, a3, a4)

	if status {
		c.JSON(200, gin.H{"status": "All good"})
	} else {
		c.JSON(500, gin.H{"There was some problem...": err})
	}
}

func apiUpdateQuestion(c *gin.Context) {
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

	status, err := dmas.UpdateQuestion(id, updateParams)

	if status {
		c.JSON(200, gin.H{"status": "All good"})
	} else {
		c.JSON(500, gin.H{"There was some problem...": err})
	}
}

func apiDeleteQuestion(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 0, 64)

	if err != nil {
		c.JSON(500, gin.H{"There was some problem...": err})
	}

	deleted := dmas.DeleteQuestion(id)
	if deleted {
		c.JSON(200, gin.H{"status": "Document is deleted"})
	} else {
		c.JSON(500, gin.H{"status": "There was some problem"})
	}

}

//function to handle get of any url - content will be returned
func apiFetch(c *gin.Context) {
	response, err := http.Get("http://" + c.Param("link"))
	if err != nil {
		c.String(400, "Your input is not valid, link should look like: www.google.com or google.com")
	} else {
		defer response.Body.Close() //close it on the end
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			c.String(500, "Specific error happend: %s", err)
			//os.Exit(1) if we want to kill the server
		} else {
			//fmt.Printf("%s\n", string(contents))
			c.String(200, string(contents))
		}
	}
}
