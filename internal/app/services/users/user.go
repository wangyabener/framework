package users

import (
	"fmt"
	"framework/database"
	"framework/internal/http/response"
	"framework/internal/models"
	"log"

	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context) {
	repository := database.NewRepostitory[models.User, models.UserEntity](database.DB)

	user := models.UserEntity{
		Username: ctx.PostForm("username"),
		Password: ctx.PostForm("password"),
		Mobile:   ctx.PostForm("mobile"),
		Email:    ctx.PostForm("email"),
	}

	err := repository.Insert(ctx, &user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)
}

func FindById(ctx *gin.Context) models.UserEntity {
	repository := database.NewRepostitory[models.User, models.UserEntity](database.DB)

	res, err := repository.FindById(ctx, ctx.Query("id"))
	if err != nil {
		log.Fatal(err)
	}

	return res
}

// func (u *UserService) Pagination(query models.UserQueryInfo) (response.Pagination[models.User], error) {
// Wrapper
// wrapper := make(map[string]interface{}, 0)
// if query.Username != "" {
// 	wrapper["username"] = query.Username
// }

// QueryBuilder
// query := database.DB.Where("username = ?", "linchong")
// }
func Pagination() (*response.Pagination[models.User], error) {
	query := database.DB.Where("1=1")

	// Paginator
	page := &database.Page[models.User]{CurrentPage: 1, PageSize: 15}

	err := page.SelectPages(query)
	if err != nil {
		fmt.Println(err)
	}

	res := response.PageResponse(page)

	return res, err
}
