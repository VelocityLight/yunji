package entity

import (
	"gorm.io/gorm"
)

type Base struct {
	gorm.Model
	IsDeleted bool `gorm:"default:false"`
}
