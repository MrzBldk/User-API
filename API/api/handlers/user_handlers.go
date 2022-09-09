package handlers

import (
	"github.com/MrzBldk/User-API/api/presenter"
	"github.com/MrzBldk/User-API/pkg/entities"
	"github.com/MrzBldk/User-API/pkg/user"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()

// AddUser is a function to add a user
// @Summary Add user
// @Description Add user
// @Tags users
// @Accept json
// @Produce json
// @Param user body presenter.User true "User"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /api/user [post]
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

// UpdateUser is a function to update a user
// @Summary Update user
// @Description Update user
// @Tags users
// @Accept json
// @Produce json
// @Param user body presenter.User true "User"
// @Success 204
// @Failure 400
// @Failure 500
// @Router /api/user [put]
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

// RemoveUser is a function to remove a user by ID
// @Summary Remove user by ID
// @Description Remove user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 204
// @Failure 404
// @Failure 500
// @Router /api/user/{id} [delete]
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

// GetUser is a function to get a user by ID
// @Summary Get user by ID
// @Description Get user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200
// @Failure 404
// @Failure 500
// @Router /api/user/{id} [get]
func GetUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := c.Params("id")
		fetched, err := service.FetchUser(userId)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.Status(fiber.StatusNotFound)
			} else {
				c.Status(fiber.StatusInternalServerError)
			}
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.GetUserSuccessResponse(fetched))
	}
}

// GetUsers is a function to get all users data from database
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200
// @Failure 500
// @Router /api/user [get]
func GetUsers(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.FetchUsers()
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.Status(fiber.StatusNotFound)
			} else {
				c.Status(fiber.StatusInternalServerError)
			}
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.GetUsersSuccessResponse(fetched))
	}
}
