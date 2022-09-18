package main

import (
	"context"

	"github.com/marattagian/inventory-system/database"
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
		),
		fx.Invoke(),
	)

	app.Run()
}
