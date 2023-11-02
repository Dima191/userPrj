package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
	"petProject/internal/app/models"
	"petProject/internal/handler"
)

const (
	userURL  = "/users/:email"
	usersURL = "/users"
	signUp   = "/users/sign-up"
	signIn   = "/users/sign-in"
)

type Handler struct {
	rep *Repository
}

func (h *Handler) Register(r *gin.Engine) {

	r.GET(usersURL, h.rep.userIdentity, h.FindAll())
	r.GET(userURL, h.FindByEmail())
	r.POST(signUp, h.SignUp())
	r.POST(signIn, h.SignIn())
}

func (h *Handler) FindAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := h.rep.FindAll(context.TODO())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, users)
	}
}

func (h *Handler) FindByEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Param("email")

		usr, err := h.rep.FindByEmail(context.TODO(), email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, usr)
	}
}

func (h *Handler) SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var usr *models.User
		if err := c.ShouldBindJSON(&usr); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		if err := usr.Validate(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		if err := usr.BeforeCreate(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		id, err := h.rep.SignUp(context.TODO(), *usr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, id)
	}
}

func (h *Handler) SignIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		usr := models.User{}
		if err := c.ShouldBindJSON(&usr); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		token, err := h.rep.SignIn(context.TODO(), usr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, token)
	}
}

func NewHandler(pool *pgxpool.Pool) handler.IHandler {
	return &Handler{
		rep: NewService(pool),
	}
}
