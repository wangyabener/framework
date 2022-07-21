package models

type User struct {
	Model
	HashID   string `json:"hashid" gorm:"comment: 用户ID"`
	Username string `json:"username" gorm:"comment: 用户名"`
	Password string `json:"password" gorm:"comment: 密码"`
	Email    string `json:"email" gorm:"comment: 邮箱"`
	Mobile   string `json:"mobile" gorm:"comment: 手机号"`
}

type UserEntity struct {
	Username string
	Password string
	Mobile   string
	Email    string
}

func (User) TableName() string {
	return "users"
}

func (u User) ToEntity() UserEntity {
	return UserEntity{
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
		Mobile:   u.Mobile,
	}
}

func (u User) FromEntity(user UserEntity) interface{} {
	return User{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Mobile:   user.Mobile,
	}
}
