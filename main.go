package main

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/marattagian/inventory-system/database"
	"github.com/marattagian/inventory-system/internal/api"
	"github.com/marattagian/inventory-system/internal/repository"
	"github.com/marattagian/inventory-system/internal/service"
	"github.com/marattagian/inventory-system/settings"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
			repository.New,
			service.New,
			api.New,
			echo.New,
		),
		fx.Invoke(
			setLifeCycle,
		),
	)

	app.Run()
}

func setLifeCycle(lc fx.Lifecycle, a *api.API, s *settings.Settings, e *echo.Echo)  {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {

			address := fmt.Sprintf(":%s", s.Port)
			go a.Start(e, address)
			
			return nil
		},
	})
}
