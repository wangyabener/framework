package user

import (
	"framework/internal/app/services/users"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	data, _ := users.Pagination()

	c.JSON(200, data)
}
