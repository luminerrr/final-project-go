package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserId uint 
	PhotoId uint 
	Message string `gorm:"not null" validator:"required~comment message cannot be empty"`
	User User `gorm:"foreignKey:UserId"`
	Photo Photo `gorm:"foreignKey:PhotoId"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)
	
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
