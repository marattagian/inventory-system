package repository

import (
	"context"

	"github.com/marattagian/inventory-system/internal/entity"
)

const (
	queryInsertUser = `
		insert into users (email, name, password)
		values ($1, $2, $3);`

	queryGetUserByEmail = `
		select id, email, name, password
		from users
		where email = $1;`

	queryInsertUserRole = `
		insert into user_roles (user_id, role_id)
		values (:user_id, :role_id);`

	queryRemoveUserRole = `
		delete from user_roles
		where user_id = :user_id and role_id = :role_id;`
)

func (r *repo) SaveUser(ctx context.Context, email, name, password string) error {
	_, err := r.db.ExecContext(ctx, queryInsertUser, email, name, password)
	return err
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	u := &entity.User{}
	err := r.db.GetContext(ctx, u, queryGetUserByEmail, email)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *repo) SaveUserRole(ctx context.Context, userID, roleID int64) error {

	data := &entity.UserRole{
		UserID: userID,
		RoleID: roleID,
	}

	_, err := r.db.NamedExecContext(ctx, queryInsertUserRole, data)

	return err
}

func (r *repo) RemoveUserRole(ctx context.Context, userID, roleID int64) error {
	data := &entity.UserRole{
		UserID: userID,
		RoleID: roleID,
	}

	_, err := r.db.NamedExecContext(ctx, queryRemoveUserRole, data)

	return err
}

func (r *repo) GetUserRoles(ctx context.Context, userID int64) ([]entity.UserRole, error) {

	roles := []entity.UserRole{}

	err := r.db.SelectContext(
		ctx,
		&roles,
		"select user_id, role_id from user_roles where user_id = $1",
		userID,
	)

	if err != nil {
		return nil, err
	}

	return roles, nil
}
