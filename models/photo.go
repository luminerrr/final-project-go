package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Title string `gorm:"not null" json:"title" valid:"required~photo title is required"`
	Caption string `gorm:"not null" json:"caption" valid:"required~photo caption is required"`
	PhotoUrl string `gorm:"not null" json:"photoUrl"`
	UserId uint 
	User User `gorm:"foreignKey:UserId"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}