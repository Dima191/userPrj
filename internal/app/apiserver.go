package app

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"petProject/internal/app/handler"
	"petProject/pkg/logger"
	"petProject/pkg/postgresql"
)

const (
	connStr = "postgres://%s:%s@%s:%s/%s"
)

type APIServer struct {
	cfg *Config
	l   *logger.Logger
}

func NewAPIServer(cfg *Config, l *logger.Logger) *APIServer {
	return &APIServer{
		cfg: cfg,
		l:   l,
	}
}

func (s *APIServer) Start() error {
	stCfg := s.cfg.Store
	pool, err := postgresql.NewClient(context.Background(), fmt.Sprintf(connStr,
		stCfg.UserName,
		stCfg.Password,
		stCfg.Host,
		stCfg.Port,
		stCfg.DBName))
	if err != nil {
		return err
	}
	defer pool.Close()

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.cfg.Host, s.cfg.Port))
	if err != nil {
		s.l.Fatalln(err)
	}

	r := gin.Default()
	h := handler.NewHandler(pool)
	h.Register(r)

	server := http.Server{
		Handler: r,
	}

	return server.Serve(listener)
}
