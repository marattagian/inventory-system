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
