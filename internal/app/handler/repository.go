package handler

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"petProject/internal/app/models"
	"petProject/internal/app/storage"
	"petProject/internal/app/storage/db/postgres"
)

type Repository struct {
	st storage.Storage
}

func (r *Repository) FindAll(ctx context.Context) (*[]models.User, error) {
	return r.st.FindAll(ctx)
}

func (r *Repository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	return r.st.FindByEmail(ctx, email)
}

func NewService(pool *pgxpool.Pool) *Repository {
	return &Repository{
		postgres.NewDB(pool),
	}
}
