package repository

import (
	"context"

	"github.com/marattagian/inventory-system/internal/entity"
)

const (
	queryInsertUser = `
insert into users (email, name, password)
values ($1, $2, $3);
`
	queryGetUserByEmail = `
select id, email, name, password
from users
where email = $1;
`
)

func (r *repo) SaveUser(ctx context.Context, email, name, password string) error {
	_, err := r.db.ExecContext(ctx, queryInsertUser, email, name, password)
	return err
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	u := &entity.User{}
	err := r.db.GetContext(ctx, u, queryGetUserByEmail, email)

	return u, err
}
