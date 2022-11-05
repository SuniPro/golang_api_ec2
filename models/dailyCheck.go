package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type DailyCheck struct {
	gorm.Model
	Username      string    `gorm:"size:70;not null" json:"username"`
	LastCheckDate time.Time `json:"last_check_date"`
}

type UserForDailyCheck struct {
	Username string
}

func (d *DailyCheck) Checking() (*DailyCheck, error) {

	var err error
	err = DB.Create(&d).Error

	if err != nil {
		return &DailyCheck{}, err
	}
	return d, nil
}

func GetUsernameByID(uid uint) (User, error) {

	var u User

	if err := DB.Select("username").First(&u, uid).Error; err != nil {
		return u, errors.New("Username not found!")
	}

	return u, nil
}

func DateEqual(lastDate, s time.Time) bool {
	y1, m1, d1 := lastDate.Date()
	y2, m2, d2 := s.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

func DateCheck(name string) (DailyCheck, error) {

	var dailyCheckData DailyCheck

	if err := DB.Last(&dailyCheckData).Select("username, last_check_date").Where("username LIKE ?", name).Error; err != nil {
		return dailyCheckData, errors.New("Username not found!")
	}
	//Model(&dailyCheckData).Select("username, last_check_date").Where("username LIKE ?", name)
	return dailyCheckData, nil
}

func (d *DailyCheck) SaveCheck() (*DailyCheck, error) {

	var err error
	err = DB.Create(&d).Error
	if err != nil {
		return &DailyCheck{}, err
	}
	return d, nil
}

//db.Select("username").Where("name LIKE ?", "name%").Table("users")

//db.Model(&User{}).Select("name, sum(age) as total").Where("id = ?", uid)
// SELECT name, sum(age) as total FROM `users` WHERE name LIKE "group%" GROUP BY `name` LIMIT 1
