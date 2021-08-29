package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-movie-project/app/entity/model"
	"github.com/xdorro/golang-fiber-movie-project/app/entity/request"
	"github.com/xdorro/golang-fiber-movie-project/app/repository"
	"github.com/xdorro/golang-fiber-movie-project/pkg/util"
)

//func FindAllUserRoles(c *fiber.Ctx) error {
//	userId := c.Params("id")
//	roles, err := repository.FindAllUserRolesByUserIdAndStatus(userId, util.STATUS_ACTIVATED)
//
//	if err != nil {
//		return util.ResponseBadRequest("ID không tồn tại", err)
//	}
//
//	result := mapper.UserRoles(*roles)
//
//	return util.ResponseSuccess("Thành công", result)
//}

// FindAllRoles : Find all roles by Status = 1
func FindAllRoles(c *fiber.Ctx) error {
	roles, err := repository.FindAllRolesByStatus(util.StatusActivated)

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", roles)
}

// FindRoleById : Find role by Role_Id and Status = 1
func FindRoleById(c *fiber.Ctx) error {
	roleId := c.Params("roleId")
	role, err := repository.FindRoleByIdAndStatus(roleId, util.StatusActivated)

	if err != nil || role.RoleId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	return util.ResponseSuccess("Thành công", role)
}

// CreateNewRole : Create a new role
func CreateNewRole(c *fiber.Ctx) error {
	roleRequest := new(request.RoleRequest)

	if err := c.BodyParser(roleRequest); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	role := model.Role{
		Name:   roleRequest.Name,
		Status: util.StatusActivated,
	}

	if _, err := repository.SaveRole(role); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

// UpdateRoleById : Update role by Role_Id and Status = 1
func UpdateRoleById(c *fiber.Ctx) error {
	roleId := c.Params("roleId")

	role, err := repository.FindRoleByIdAndStatus(roleId, util.StatusActivated)

	if err != nil || role.RoleId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	roleRequest := new(request.RoleRequest)
	if err = c.BodyParser(roleRequest); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	role.Name = roleRequest.Name

	if _, err = repository.SaveRole(*role); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

// DeleteRoleById : Delete role by Role_Id and Status = 1
func DeleteRoleById(c *fiber.Ctx) error {
	roleId := c.Params("roleId")
	role, err := repository.FindRoleByIdAndStatus(roleId, util.StatusActivated)

	if err != nil || role.RoleId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	role.Status = util.StatusDeleted

	if _, err = repository.SaveRole(*role); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}
