package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nlabsoft_assignment2/models"
	"nlabsoft_assignment2/utils/token"
)

func CountingCheck(c *gin.Context) {
	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errortoken": err.Error()})
		return
	}

	u, err := models.GetUserByID(user_id)

	d := u.Username

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorget": err.Error()})
		return
	}

	var checkDb []models.DailyCheck
	var checkCount int
	models.DB.Where("username LiKE ?", d).Find(&checkDb).Count(&checkCount)
	//DB.Last(&dailyCheckData).Select("username, last_check_date").Where("username LIKE ?", name)

	c.JSON(http.StatusOK, gin.H{"message": checkCount})

}
