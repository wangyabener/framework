package user

import (
	"framework/internal/app/services/users"

	"github.com/gin-gonic/gin"
)

type detailResponse struct {
	Id        int    `json:"id"`         // ID
	HashID    string `json:"hashid"`     // hashid
	Username  string `json:"username"`   // 用户名
	Nickname  string `json:"nickname"`   // 昵称
	Mobile    string `json:"mobile"`     // 手机号
	CreatedAt string `json:"created_at"` // 创建时间
	UpdatedAt string `json:"updated_at"` // 更新时间
}

type detailRequest struct {
	Page     int    `form:"page"`
	Size     int    `form:"size"`
	Username string `form:"username"`
}

func Detail(c *gin.Context) {
	data := users.FindById(c)

	c.JSON(200, data)
}
