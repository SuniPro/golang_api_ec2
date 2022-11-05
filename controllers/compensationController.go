package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nlabsoft_assignment2/models"
	"nlabsoft_assignment2/utils/token"
)

func CompensationController(c *gin.Context) {
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

	compensation := models.Compensation{}

	var checkDb []models.DailyCheck
	var checkCount int
	models.DB.Where("username LiKE ?", d).Find(&checkDb).Count(&checkCount)

	var firstReward string = "Adventure_pack1"
	var secondReward string = "unique_weapon"
	var thirdReward string = "pat"
	var fourthReward string = "riding_pat"
	var fifthReward string = "cash"
	var sixthReward string = "elixir"

	switch {
	case checkCount == 7:
		compensation.Username = u.Username
		compensation.Compensation = firstReward
		_, err = compensation.SaveCompensation()
		c.JSON(http.StatusOK, gin.H{"data": firstReward})
	case checkCount == 14:
		compensation.Username = u.Username
		compensation.Compensation = secondReward
		_, err = compensation.SaveCompensation()
		c.JSON(http.StatusOK, gin.H{"data": secondReward})
	case checkCount == 21:
		compensation.Username = u.Username
		compensation.Compensation = thirdReward
		_, err = compensation.SaveCompensation()
		c.JSON(http.StatusOK, gin.H{"data": thirdReward})
	case checkCount == 28:
		compensation.Username = u.Username
		compensation.Compensation = fourthReward
		_, err = compensation.SaveCompensation()
		c.JSON(http.StatusOK, gin.H{"data": fourthReward})
	case checkCount == 35:
		compensation.Username = u.Username
		compensation.Compensation = fifthReward
		_, err = compensation.SaveCompensation()
		c.JSON(http.StatusOK, gin.H{"data": fifthReward})
	case checkCount == 42:
		compensation.Username = u.Username
		compensation.Compensation = sixthReward
		_, err = compensation.SaveCompensation()
		c.JSON(http.StatusOK, gin.H{"data": sixthReward})
	default:
		c.JSON(http.StatusOK, gin.H{"data": "현재 받을 수 있는 보상이 없습니다 :)"})
	}

}
