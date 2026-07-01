package server

import (
	"database/sql"
	"bookmarks/internal/handler"
	"bookmarks/internal/middleware"

	"github.com/gin-gonic/gin"
)

type Server struct {
	db     *sql.DB
	router *gin.Engine
	port   string
}

func New(port string, db *sql.DB) *Server {
	s := &Server{db: db, router: gin.Default(), port: port}
	s.setupRoutes()
	return s
}

func (s *Server) Run() error {
	s.router.Run(s.port)
	return nil
}

func (s *Server) setupRoutes() {
	s.router.GET("/", handler.HealthHandler)

	api := s.router.Group("/api/v1")
	s.registerUserRoutes(api)

  protected := api.Group("")
	protected.Use(middleware.AuthMiddleware())
	s.registerUserProtectedRoutes(protected)
  s.registerBookmarkProtectedRoutes(protected)
}