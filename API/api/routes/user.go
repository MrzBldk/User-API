package routes

import (
	"github.com/MrzBldk/User-API/api/handlers"
	"github.com/MrzBldk/User-API/pkg/user"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, service user.Service) {
	app.Get("api/user", handlers.GetUsers(service))
	app.Get("api/user/:id", handlers.GetUser(service))
	app.Post("api/user", handlers.AddUser(service))
	app.Put("api/user", handlers.UpdateUser(service))
	app.Delete("api/user/:id", handlers.RemoveUser(service))
}
