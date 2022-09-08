package handlers

import (
	"github.com/MrzBldk/User-API/api/presenter"
	"github.com/MrzBldk/User-API/pkg/entities"
	"github.com/MrzBldk/User-API/pkg/user"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func AddUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.User
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		if err = validate.Struct(&requestBody); err != nil {
			return c.JSON(presenter.UserErrorResponse(err))
		}
		result, err := service.InsertUser(&requestBody)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		c.Status(fiber.StatusOK)
		return c.JSON(presenter.CreateUserSuccessResponse(result))
	}
}

func UpdateUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.User
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		if err = validate.Struct(&requestBody); err != nil {
			return c.JSON(presenter.UserErrorResponse(err))
		}
		err = service.UpdateUser(&requestBody)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.SendStatus(fiber.StatusNoContent)
	}
}

func RemoveUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := c.Params("id")
		err := service.RemoveUser(userId)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.SendStatus(fiber.StatusNoContent)
	}
}

func GetUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := c.Params("id")
		fetched, err := service.FetchUser(userId)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.GetUserSuccessResponse(fetched))
	}
}

func GetUsers(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.FetchUsers()
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.GetUsersSuccessResponse(fetched))
	}
}
