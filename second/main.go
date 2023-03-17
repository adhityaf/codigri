package main

import (
	"go-second-assessment/controllers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	route := gin.Default()
	mainRouter := route.Group("/v1")
	{
		mainRouter.POST("/login", controllers.Login)
		mainRouter.POST("/register", controllers.Register)

		authorized := mainRouter.Group("/")
		// authorized.Use(middlewares.Auth())
		// {
			authorized.POST("/refresh", controllers.Refresh)
			authorized.GET("/profile", controllers.Profile)
		
			authorized.GET("/articles", controllers.GetArticles)
			authorized.GET("/articles/:id", controllers.GetOneArticle)
		// }		
	}

	route.Run(":8000")
}
