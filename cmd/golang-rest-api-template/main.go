package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/MykhayloGusak/golang-rest-api-template/config"
	usersController "github.com/MykhayloGusak/golang-rest-api-template/internal/app/golang-rest-api-template/users/controller"
	usersRepository "github.com/MykhayloGusak/golang-rest-api-template/internal/app/golang-rest-api-template/users/repository"
	usersService "github.com/MykhayloGusak/golang-rest-api-template/internal/app/golang-rest-api-template/users/service"
	"github.com/MykhayloGusak/golang-rest-api-template/internal/pkg/mongodb"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// create a context for the server and a function to stop it
	serverCtx, serverCancelCtx := context.WithCancel(context.Background())
	defer serverCancelCtx()

	// load configuration
	cfg := config.Get()
	// log.Println("Configuration loaded:", cfg) // TODO: hide from logs

	// initialize MongoDB
	mdb, err := mongodb.New(serverCtx, cfg.MongodbURI, cfg.MongodbDatabase)
	if err != nil {
		log.Printf("Failed to connect to MongoDB: %v\n", err)
		return err
	}
	defer mdb.Disconnect(serverCtx)
	log.Println("MongoDB connected")

	// initialize user repository and service
	userRepo := usersRepository.New(mdb.UsersCollection)
	userService := usersService.New(userRepo)
	defer userService.Cancel()

	// initialize user controller
	userController := usersController.New(userService)

	// default config
	api := fiber.New()

	// add middleware
	api.Use(cors.New())

	// Register HTTP routes using Fiber app
	userController.RegisterRoutes(api.Group("/api/users"))

	// add health check
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Healthy!")
	})

	go func() {
		shutdown := make(chan os.Signal, 1)
		signal.Notify(shutdown, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

		<-shutdown
		log.Println("Shutdown signal received..")
		serverCancelCtx()

		if err := api.Shutdown(); err != nil {
			log.Printf("Failed to shutdown the server: %v\n", err)
		}
	}()

	// run the server
	if err := api.Listen(fmt.Sprintf(":%s", cfg.HTTPPort)); err != nil {
		return fmt.Errorf("server failed to start: %w", err)
	}

	log.Println("Server stopped")

	os.Exit(0)

	return nil
}
