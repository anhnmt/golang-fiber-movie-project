package mapper

import (
	model2 "github.com/xdorro/golang-fiber-base-project/app/entity/model"
)

func UserRoles(roles []model2.Role) []model2.Role {
	result := make([]model2.Role, 0)

	for _, role := range roles {
		if role.RoleId != 0 {
			result = append(result, role)
		}
	}
	return result
}
