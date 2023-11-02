package storage

import (
	"context"
	"petProject/internal/app/models"
)

type Storage interface {
	FindAll(ctx context.Context) (*[]models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	Create(ctx context.Context, user models.User) (int, error)
}
