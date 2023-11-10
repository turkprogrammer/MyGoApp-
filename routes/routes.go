package routes

import (
	"github.com/gin-gonic/gin"
	"mygoapp/handlers"
)

func SetupRoutes() {
	router := gin.Default()

	// Пример маршрутов
	router.GET("/hello", handlers.HelloWorld)
	router.GET("/show-page", handlers.ShowPage)
	router.POST("/create-user", handlers.CreateUser)
	router.GET("/get-users", handlers.GetUsers)

	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	router.Run(":8080")
}
