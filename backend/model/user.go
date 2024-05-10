package model

import "gorm.io/gorm"

type User struct{
	gorm.Model
	Name     string `json:"name"`
    Email    string `gorm:"unique" json:"email"`
    Password string `json:"password"`
    Role     string `json:"role"`
}
type CreateFeedBack struct {

}
type EditFeedBack struct{

}

type PostComment struct{

}