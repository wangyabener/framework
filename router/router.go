package router

import (
	"framework/internal/app/api/v1/user"

	"github.com/gin-gonic/gin"
)

func NewHttpServer() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/users/:id", user.Detail)
		v1.GET("/users", user.List)
	}

	return router
}
