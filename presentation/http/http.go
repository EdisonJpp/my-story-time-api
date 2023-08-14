package http

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/fx"
)

func ProvideHttp(lc fx.Lifecycle) *fiber.App {
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to the authentication-services!"))
	})

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				err := app.Listen(":5000")

				if err != nil {
					fmt.Printf("Error starting server: %v\n", err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})

	return app
}
