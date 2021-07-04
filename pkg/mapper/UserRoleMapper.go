package mapper

import (
	"github.com/xdorro/golang-fiber-base-project/app/model"
)

func UserRoles(roles []model.Role) []model.Role {
	result := make([]model.Role, 0)

	for _, role := range roles {
		if role.RoleId != 0 {
			result = append(result, role)
		}
	}
	return result
}
