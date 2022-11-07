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
	r.Use(CORSMiddleware())

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)
	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)
	protected.POST("/daily_check", controllers.Checker)
	protected.GET("/compensation", controllers.CompensationController)
	protected.GET("/counting_check", controllers.CountingCheck)

	r.Run(":8080")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, DELETE, POST, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
