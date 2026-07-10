package server

import (
	"bookmarks/internal/handler"
	"bookmarks/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/getsentry/sentry-go/gin"
)

type Server struct {
	db     *gorm.DB
	router *gin.Engine
	port   string
}

func New(port string, db *gorm.DB) *Server {
	s := &Server{db: db, router: gin.New(), port: port}
  
	s.router.Use(sentrygin.New(sentrygin.Options{
		Repanic: true, // чтобы следующий recovery тоже сработал
	}))
	s.router.Use(gin.Recovery())              // panic → 500
	s.router.Use(middleware.RequestID())
	s.router.Use(middleware.RequestLogger())

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
	s.registerAuthRoutes(api)

  protected := api.Group("")
	protected.Use(middleware.AuthMiddleware())
	s.registerUserRoutes(protected)
  s.registerBookmarkProtectedRoutes(protected)
  s.registerTagProtectedRoutes(protected)
}