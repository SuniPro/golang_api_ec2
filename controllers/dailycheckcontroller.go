package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"nlabsoft_assignment2/models"
	"nlabsoft_assignment2/utils/token"
	"time"
)

func Checker(c *gin.Context) {

	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := models.GetUsernameByID(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	d := models.DailyCheck{}
	s := time.Now()

	d.Username = u.Username
	d.LastCheckDate = s

	lastCheckDate, err := models.DateCheck(d.Username)

	lastDate := lastCheckDate.LastCheckDate

	if models.DateEqual(lastDate, s) == false {
		fmt.Print(models.DateEqual(lastDate, s))
		_, err = d.SaveCheck()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "출석했습니다 !"})

	} else {
		c.JSON(http.StatusOK, gin.H{"message": "이미 출석하였습니다"})
	}

	//c.JSON(http.StatusOK, gin.H{"message": "success", "data": d})

}
