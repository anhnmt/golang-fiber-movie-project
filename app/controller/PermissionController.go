package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/app/entity/request"
	"github.com/xdorro/golang-fiber-base-project/app/repository"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
)

// FindAllPermissions : Find all permissions by Status = 1
func FindAllPermissions(c *fiber.Ctx) error {
	permissions, err := repository.FindAllPermissionsByStatus(util.StatusActivated)

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", permissions)
}

// FindPermissionById : Find permission by Permission_Id and Status = 1
func FindPermissionById(c *fiber.Ctx) error {
	permissionId := c.Params("id")
	permission, err := repository.FindPermissionByIdAndStatus(permissionId, util.StatusActivated)

	if err != nil || permission.PermissionId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	return util.ResponseSuccess("Thành công", permission)
}

// CreateNewPermission : Create a new permission
func CreateNewPermission(c *fiber.Ctx) error {
	permissionRequest := new(request.PermissionRequest)

	if err := c.BodyParser(permissionRequest); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	permission := model.Permission{
		Name:   permissionRequest.Name,
		Status: util.StatusActivated,
	}

	if _, err := repository.SavePermission(permission); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

// UpdatePermissionById : Update permission by Permission_Id and Status = 1
func UpdatePermissionById(c *fiber.Ctx) error {
	permissionId := c.Params("id")

	permission, err := repository.FindPermissionByIdAndStatus(permissionId, util.StatusActivated)

	if err != nil || permission.PermissionId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	permissionRequest := new(request.PermissionRequest)
	if err = c.BodyParser(permissionRequest); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	permission.Name = permissionRequest.Name

	if _, err = repository.SavePermission(*permission); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

// DeletePermissionById : Delete permission by Permission_Id and Status = 1
func DeletePermissionById(c *fiber.Ctx) error {
	permissionId := c.Params("id")
	permission, err := repository.FindPermissionByIdAndStatus(permissionId, util.StatusActivated)

	if err != nil || permission.PermissionId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	permission.Status = util.StatusDeleted

	if _, err = repository.SavePermission(*permission); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}
