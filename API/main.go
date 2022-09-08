package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/MrzBldk/User-API/api/routes"
	"github.com/MrzBldk/User-API/pkg/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file, %v", err)
	}

	logfile, err := os.Create("file.log")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer logfile.Close()

	db, cancel, err := databaseConnection()
	if err != nil {
		log.Fatalf("Database Connection Error %v", err)
	}

	userCollection := db.Collection("users")
	userRepo := user.NewRepo(userCollection)
	userService := user.NewService(userRepo)

	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Output: logfile,
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n${resBody}\n\n",
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, X-Requested-With, Content-Type, Accept",
		AllowMethods: "GET, PUT, POST, DELETE",
	}))

	routes.UserRouter(app, userService)

	defer cancel()

	log.Fatal(app.Listen(":8080"))
}

func databaseConnection() (*mongo.Database, context.CancelFunc, error) {
	uri := os.Getenv("MONGOURI")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		cancel()
		return nil, nil, err
	}
	db := client.Database("usersdb")
	return db, cancel, nil
}
