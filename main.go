package main

import (
	"github.com/gin-gonic/gin"
	"nlabsoft_assignment2/controllers"
	"nlabsoft_assignment2/middlewares"
	"nlabsoft_assignment2/models"
)

func main() {

	models.ConnectDataBase()

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/", controllers.Print)
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)
	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)
	protected.GET("/daily_check", controllers.Checker)
	protected.GET("/compensation", controllers.CompensationController)
	protected.GET("/counting_check", controllers.CountingCheck)

	r.Run(":8080")
}
