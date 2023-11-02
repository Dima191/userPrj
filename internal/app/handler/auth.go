package handler

import (
	"context"
	"database/sql"
	"errors"
	"github.com/golang-jwt/jwt"
	errors2 "petProject/internal/app/errors"
	"petProject/internal/app/models"
	"time"
)

const (
	signingKey = "5iwbzI9YXH920MVJjUEaAPuiw7Cl0cLLetYoOQMlHTwNenMmn0"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId string
}

func (r *Repository) SignUp(ctx context.Context, user models.User) (int, error) {
	return r.st.Create(ctx, user)
}

func (r *Repository) SignIn(ctx context.Context, user models.User) (string, error) {
	token, err := r.generateToken(ctx, user.Email, user.Password)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *Repository) generateToken(ctx context.Context, email, password string) (string, error) {
	usr, err := r.st.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errors2.NotFound
		}
		return "", err
	}

	if err = models.ComparePasswords(usr.HashPassword, password); err != nil {
		return "", errors2.WrongEmailPassword
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: usr.ID,
	})

	return token.SignedString([]byte(signingKey))
}

func (r *Repository) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors2.InvalidSigningMethod
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors2.TokenClaimsErr
	}

	return claims.UserId, nil
}
