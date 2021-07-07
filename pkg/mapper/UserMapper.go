package mapper

import (
	model2 "github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/app/entity/response"
)

func ListUserSearch(users []model2.User) []response.UserResponse {
	result := make([]response.UserResponse, 0)

	for _, user := range users {
		mapper := UserSearch(&user)
		result = append(result, mapper)
	}
	return result
}

func UserSearch(user *model2.User) response.UserResponse {
	return response.UserResponse{
		UserId:   user.UserId,
		Name:     user.Name,
		Username: user.Username,
		Gender:   user.Gender,
	}
}
