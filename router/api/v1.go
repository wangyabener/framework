package api

import (
	"framework/internal/app/services/user"
	"framework/router"
)

func Register() {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/users", user.List)
	}
}
