package service

import (
	"context"
	"errors"

	"github.com/marattagian/inventory-system/encryption"
	"github.com/marattagian/inventory-system/internal/models"
)

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid password")
	ErrRoleAlreadyAdded   = errors.New("role was already added to user")
	ErrRoleNotFound       = errors.New("role not found")
)

func (s *serv) RegisterUser(ctx context.Context, email, name, password string) error {
	u, _ := s.repo.GetUserByEmail(ctx, email)
	if u != nil {
		return ErrUserAlreadyExists
	}

	cipheredBytes, err := encryption.Encrypt([]byte(password))
	if err != nil {
		return err
	}

	cipheredPassword := encryption.ToBase64(cipheredBytes)
	return s.repo.SaveUser(ctx, email, name, cipheredPassword)
}

func (s *serv) LoginUser(ctx context.Context, email, password string) (*models.User, error) {
	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	hashedBytes, err := encryption.FromBase64(u.Password)
	if err != nil {
		return nil, err
	}

	decryptedPassword, err := encryption.Decrypt(hashedBytes)
	if err != nil {
		return nil, err
	}

	if u.Password != string(decryptedPassword) {
		return nil, ErrInvalidCredentials
	}

	return &models.User{
		ID:    u.ID,
		Email: u.Email,
		Name:  u.Name,
	}, nil
}

func (s *serv) AddUserRole(ctx context.Context, userID, roleID int64) error {

	roles, err := s.repo.GetUserRoles(ctx, userID)
	if err != nil {
		return err
	}

	for _, r := range roles {
		if r.RoleID == roleID {
			return ErrRoleAlreadyAdded
		}
	}

	return s.repo.SaveUserRole(ctx, userID, roleID)
}

func (s *serv) RemoveUserRole(ctx context.Context, userID, roleID int64) error {

	roles, err := s.repo.GetUserRoles(ctx, userID)
	if err != nil {
		return err
	}

	for _, r := range roles {
		if r.RoleID == roleID {
			return s.repo.RemoveUserRole(ctx, userID, roleID)
		}
	}

	return ErrRoleNotFound
}
