package mapper

import (
	model "github.com/xdorro/golang-fiber-movie-project/app/entity/model"
	"github.com/xdorro/golang-fiber-movie-project/app/entity/response"
)

func ListUserSearch(users []model.User) []response.UserResponse {
	result := make([]response.UserResponse, 0)

	for _, user := range users {
		mapper := UserSearch(&user)
		result = append(result, mapper)
	}
	return result
}

func UserSearch(user *model.User) response.UserResponse {
	return response.UserResponse{
		UserId:   user.UserId,
		Name:     user.Name,
		Username: user.Username,
		Gender:   user.Gender,
	}
}
