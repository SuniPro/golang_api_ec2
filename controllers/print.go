package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Print(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "출력"})
}
