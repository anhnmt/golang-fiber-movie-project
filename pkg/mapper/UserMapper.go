package mapper

import (
	"github.com/xdorro/golang-fiber-base-project/app/dto"
	"github.com/xdorro/golang-fiber-base-project/app/model"
)

func ListUserSearch(users []model.User) []dto.SearchUserResponse {
	var result []dto.SearchUserResponse

	for _, user := range users {
		mapper := UserSearch(&user)
		result = append(result, mapper)
	}
	return result
}

func UserSearch(user *model.User) dto.SearchUserResponse {
	return dto.SearchUserResponse{
		UserId:   user.UserId,
		Name:     user.Name,
		Username: user.Username,
		Gender:   user.Gender,
	}
}
