package user

import (
	"github.com/gin-gonic/gin"
)

type listItem struct {
	Id        int    `json:"id"`         // ID
	HashID    string `json:"hashid"`     // hashid
	Username  string `json:"username"`   // 用户名
	Nickname  string `json:"nickname"`   // 昵称
	Mobile    string `json:"mobile"`     // 手机号
	CreatedAt string `json:"created_at"` // 创建时间
	UpdatedAt string `json:"updated_at"` // 更新时间
}

type listRequest struct {
	Page     int    `form:"page"`
	Size     int    `form:"size"`
	Username string `form:"username"`
}

type listResponse struct {
	List       []listItem `json:"list"`
	Pagination struct {
		Total       int `json:"total"`
		CurrentPage int `json:"current_page"`
		PerPage     int `json:"per_page"`
	}
}

func List(c *gin.Context) {

}
