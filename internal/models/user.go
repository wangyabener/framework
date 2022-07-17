package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"comment: 用户名"`
	Password string `json:"password" gorm:"comment: 密码"`
	Email    string `json:"email" gorm:"comment: 邮箱"`
	Mobile   string `json:"mobile" gorm:"comment: 手机号"`
}

func (User) TableName() string {
	return "users"
}
