package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/repository"
	"github.com/xdorro/golang-fiber-base-project/pkg/mapper"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
)

func FindAllUserRoles(c *fiber.Ctx) error {
	userId := c.Params("id")
	roles, err := repository.FindAllUserRolesByUserIdAndStatus(userId, 1)

	if err != nil {
		return util.ResponseBadRequest(c, "ID không tồn tại", err)
	}

	result := mapper.UserRoles(*roles)

	return util.ResponseSuccess(c, "Thành công", result)
}
