package database

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/marattagian/inventory-system/settings"
)

func New(ctx context.Context, s *settings.Settings) (*sqlx.DB, error) {
	connectionString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		s.DB.User,
		s.DB.Password,
		s.DB.Host,
		s.DB.Port,
		s.DB.Name,
	)

	return sqlx.ConnectContext(ctx, "postgres", connectionString)
}
