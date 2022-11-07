package models

import (
	"github.com/jinzhu/gorm"
)

type Compensation struct {
	gorm.Model
	Email        string `gorm:"size:70;not null" json:"email"`
	Username     string `gorm:"size:70;not null" json:"username"`
	Compensation string `gorm:"size:70;not null" json:"username"`
}

func (compensation *Compensation) SaveCompensation() (*Compensation, error) {

	var err error
	err = DB.Create(&compensation).Error
	if err != nil {
		return &Compensation{}, err
	}
	return compensation, nil
}
