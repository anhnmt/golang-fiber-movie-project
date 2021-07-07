package controller

import (
	"github.com/gofiber/fiber/v2"
	model2 "github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/app/entity/request"
	"github.com/xdorro/golang-fiber-base-project/app/repository"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
)

//func FindAllUserRoles(c *fiber.Ctx) error {
//	userId := c.Params("id")
//	roles, err := repository.FindAllUserRolesByUserIdAndStatus(userId, util.STATUS_ACTIVATED)
//
//	if err != nil {
//		return util.ResponseBadRequest(c, "ID không tồn tại", err)
//	}
//
//	result := mapper.UserRoles(*roles)
//
//	return util.ResponseSuccess(c, "Thành công", result)
//}

// FindAllRoles : Find all roles by Status = 1
func FindAllRoles(c *fiber.Ctx) error {
	roles, err := repository.FindAllRolesByStatus(util.STATUS_ACTIVATED)

	if err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", roles)
}

// FindRoleById : Find role by Role_Id and Status = 1
func FindRoleById(c *fiber.Ctx) error {
	roleId := c.Params("id")
	role, err := repository.FindRoleByIdAndStatus(roleId, util.STATUS_ACTIVATED)

	if err != nil || role.RoleId == 0 {
		return util.ResponseBadRequest(c, "ID không tồn tại", err)
	}

	return util.ResponseSuccess(c, "Thành công", role)
}

// CreateNewRole : Create a new role
func CreateNewRole(c *fiber.Ctx) error {
	roleRequest := new(request.RoleRequest)

	if err := c.BodyParser(roleRequest); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	role := model2.Role{
		Name:   roleRequest.Name,
		Status: util.STATUS_ACTIVATED,
	}

	if _, err := repository.SaveRole(role); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}

// UpdateRoleById : Update role by Role_Id and Status = 1
func UpdateRoleById(c *fiber.Ctx) error {
	roleId := c.Params("id")

	role, err := repository.FindRoleByIdAndStatus(roleId, util.STATUS_ACTIVATED)

	if err != nil || role.RoleId == 0 {
		return util.ResponseBadRequest(c, "ID không tồn tại", err)
	}

	roleRequest := new(request.RoleRequest)
	if err = c.BodyParser(roleRequest); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	role.Name = roleRequest.Name

	if _, err = repository.SaveRole(*role); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}

// DeleteRoleById : Delete role by Role_Id and Status = 1
func DeleteRoleById(c *fiber.Ctx) error {
	roleId := c.Params("id")
	role, err := repository.FindRoleByIdAndStatus(roleId, util.STATUS_ACTIVATED)

	if err != nil || role.RoleId == 0 {
		return util.ResponseBadRequest(c, "ID không tồn tại", err)
	}

	role.Status = util.STATUS_DELETED

	if _, err = repository.SaveRole(*role); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}
