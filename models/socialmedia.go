package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	gorm.Model
	Name string `gorm:"not null" json:"name" validator:"required~social media name is required"`
	SocialMediaUrl string `gorm:"not null" json:"socialMediaUrl" validator:"required~social media url is required"`
	UserId uint 
	User User `gorm:"foreignKey:UserId"`
}

func (sc *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(sc)
	
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}