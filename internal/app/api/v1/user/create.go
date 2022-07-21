package user

import (
	"framework/internal/app/services/users"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	users.Create(c)
}
