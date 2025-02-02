package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Ilham-muttaqien17/learn-restful-go/config"
	"github.com/Ilham-muttaqien17/learn-restful-go/middlewares"
	"github.com/Ilham-muttaqien17/learn-restful-go/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load env file
	if err := config.LoadEnv(); err != nil {
		fmt.Println("❌ Error loading env file: ", err)
		os.Exit(1)
	}

	// Initialize redis connection
	if err := config.RegisterRedis(); err != nil {
		fmt.Println("❌ Error connecting redis storage: ", err)
		os.Exit(1)
	}
	
	fmt.Println("✅ Redis connection established")

	// Initialize database connection
	if err := config.ConnectDB(); err != nil {
		fmt.Println("❌ Error connecting database: ", err)
		os.Exit(1)
	}

	fmt.Println("✅ Database connection opened")

	// Initialize Fiber app
	app := fiber.New()

	// Register middlewares
	middlewares.Register(app)

	// Register routes
	routes.Register(app)

	// Handle Unmatched Routes
	app.Use("*", func (ctx *fiber.Ctx) error  {
		return ctx.Status(fiber.StatusNotFound).JSON(map[string]interface{}{
			"message": fmt.Sprintf("Route %s:%s not found", ctx.Method(),ctx.Path()),
		})
	})

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	address := fmt.Sprintf("%s:%d", config.Env.AppHost, config.Env.AppPort)

	// Start the server in a goroutine
	go func() {
		if err := app.Listen(address); err != nil {
			fmt.Println("❌ Error starting server: ", err)
			os.Exit(1)
		}
	}()

	// Block until we receive a termination signal
	<-quit
	fmt.Println("⌚ Shutting down gracefully...")

	if err := app.Shutdown(); err != nil {
		fmt.Println("❌ Error during shutdown: ", err)
	}

	if err := config.CloseRedis(); err != nil {
		fmt.Println("❌ Error during close redis connection: ", err)
	}

	if err := config.DisconnectDB(); err != nil {
		fmt.Println("❌ Error during close database connection: ", err)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("✅ Redis connection closed")
	fmt.Println("✅ Database connection closed")
	fmt.Println("✅ Server successfully shutting down")

}