package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email string `gorm:"size:63;not null;unique;index"`
	Password string `gorm:"size:63;not null"`
	Avatar string `gorm:"size:63"`
	Nickname string `gorm:"size:63;index"`
	IsActivate bool `gorm:"default:false;index"`
}
