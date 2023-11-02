package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
	errors2 "petProject/internal/app/errors"
	"petProject/internal/app/models"
	"petProject/internal/app/storage"
)

type DB struct {
	*pgxpool.Pool
}

func (d *DB) FindAll(ctx context.Context) (*[]models.User, error) {
	q := "select id,name,email,hash_password from users"
	rows, err := d.Query(ctx, q)
	defer rows.Close()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors2.NotFound
		}
		return nil, err
	}

	users := make([]models.User, 0)

	for rows.Next() {
		var u models.User
		err = rows.Scan(&u.ID, &u.Name, &u.Email, &u.HashPassword)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &users, nil
}

func (d *DB) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	q := "select id,name,email,hash_password from users where email = $1"
	row := d.QueryRow(ctx, q, email)
	u := models.User{}
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.HashPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors2.NotFound
		}
		return nil, err
	}
	return &u, nil
}

func (d *DB) Create(ctx context.Context, user models.User) (int, error) {
	q := "insert into users (name, email, hash_password) values ($1, $2, $3) RETURNING id"

	row := d.QueryRow(ctx, q, user.Name, user.Email, user.HashPassword)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func NewDB(pool *pgxpool.Pool) storage.Storage {
	return &DB{
		pool,
	}
}
