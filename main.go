package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Ilham-muttaqien17/learn-restful-go/models"
	"github.com/Ilham-muttaqien17/learn-restful-go/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	if err := models.ConnectDB(); err != nil {
		fmt.Println("Error connecting database: ", err)
	}

	app := fiber.New()

	/* Register routes */
	routes.Register(app)

	/* Handle Unmatched Routes */
	app.Use("*", func (ctx *fiber.Ctx) error  {
		return ctx.Status(fiber.StatusNotFound).JSON(map[string]interface{}{
			"message": fmt.Sprintf("Route %s:%s not found", ctx.Method(),ctx.Path()),
		})
	})

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	/* Start the server in a goroutine */
	go func() {
		if err := app.Listen(":3300"); err != nil {
			fmt.Println("Error starting server: ", err)
		}
	}()

	// Block until we receive a termination signal
	<-quit
	fmt.Println("Shutting down gracefully...")

	if err := app.Shutdown(); err != nil {
		fmt.Println("Error during shutdown: ", err)
	}

	if err := models.DisconnectDB(); err != nil {
		fmt.Println("Error during close database connection: ", err)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Database connection closed")
	fmt.Println("Server successfully shutting down")

}