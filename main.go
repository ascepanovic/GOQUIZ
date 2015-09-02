package main

import (
	"dmas/controllers"
	"dmas/utils/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {

	port := ":4555" //on this port our web server will be liste for incoming requests
	r := gin.Default() //default instance of gin framework used as router as well

	//CORSMiddlware will inject needed headers for our angular client
	r.Use(middlewares.CORS())

	//Homepage
	r.GET("/", controllers.Home)

	//WebSocket
	r.GET("/ws", func(c *gin.Context) {
		controllers.Wshandler(c.Writer, c.Request)
	})

	// Now we will do some routes grouping and there we are going to have basic post/get functions/handlers
	userRoutes := r.Group("/user")
	{
		userRoutes.POST("/register", controllers.RegisterHandler)
		userRoutes.GET("/login", controllers.LoginHandler)
	}

	// Grouped API routes
	apiRoutes := r.Group("api")
	{
		// API routes for users
		apiRoutes.GET("/users", controllers.GetAllUsers)
		apiRoutes.GET("/user/:id", controllers.GetOneUser)
		apiRoutes.POST("/user/create", controllers.CreateUser)
		apiRoutes.POST("/user/update", controllers.UpdateUser)
		apiRoutes.POST("/user/delete/:id", controllers.DeleteUser)

		// API routes for questions
		apiRoutes.GET("/questions", controllers.GetAllQuestions)
		apiRoutes.GET("/question/:id", controllers.GetOneQuestion)
		apiRoutes.POST("/question/create", controllers.CreateQuestion)
		apiRoutes.POST("/question/update", controllers.UpdateQuestion)
		apiRoutes.POST("/question/delete/:id", controllers.DeleteQuestion)
	}

	//now listen on defined port
	r.Run(port)
}
