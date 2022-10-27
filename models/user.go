package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
	"final-project-go/helpers"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null" valid:"required~Your username is required" json:"username"`
	Email string `gorm:"not null;uniqueIndex" valid:"required~your email is required" json:"email"`
	Password string `gorm:"not null" valid:"required~your password is required, minstringlength(6)~password has to have a minimum length of 6" json:"password"`
	Age int `json:"age"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	
	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}