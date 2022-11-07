package controllers

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"nlabsoft_assignment2/models"
	"nlabsoft_assignment2/utils/token"
)

type RegisterInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {

	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	var nick = [10]string{"이상한 수니", "신기한 수니", "신비한 수니", "놀라운 수니", "굉장한 수니", "똑똑한 수니", "영리한 수니", "어여쁜 수니", "매력있는 수니", "재밌는 수니"}
	//var name = [10]string{"수니", "가니", "여니", "주니", "미니", "우니", "혀니", "아니", "무니", "사니"}

	u.Email = input.Email
	u.Username = nick[rand.Intn(9)]
	u.Password = input.Password

	_, err := u.SaveUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})
}

func Login(c *gin.Context) {

	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Email = input.Email
	u.Password = input.Password

	token, err := models.LoginCheck(u.Email, u.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

func CurrentUser(c *gin.Context) {

	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := models.GetUserByID(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": u.Email, "data": u.Username})

}
