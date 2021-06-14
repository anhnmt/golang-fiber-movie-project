package mapper

import (
	"github.com/xdorro/golang-fiber-base-project/app/dto"
	"github.com/xdorro/golang-fiber-base-project/app/model"
)

func ListUserSearch(users []model.User) []dto.UserResponse {
	var result []dto.UserResponse

	for _, user := range users {
		mapper := UserSearch(&user)
		result = append(result, mapper)
	}
	return result
}

func UserSearch(user *model.User) dto.UserResponse {
	return dto.UserResponse{
		UserId:   user.UserId,
		Name:     user.Name,
		Username: user.Username,
		Gender:   user.Gender,
	}
}
