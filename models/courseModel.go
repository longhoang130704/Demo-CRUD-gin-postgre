package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Id int `gorm:"primaryKey"`
	Title string `form:"title" binding:"required"`
	Description string `form:"description" binding:"required"`
	ImgLink string `form:"linkImg" binding:"required"`
	VideoLink string `form:"linkVideo" binding:"required"`

}
