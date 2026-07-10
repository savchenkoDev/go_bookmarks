package server

import (
	"bookmarks/internal/cache"
	"bookmarks/internal/config"
	"bookmarks/internal/handler"
	"bookmarks/internal/middleware"

	"github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	db     *gorm.DB
	cache  *cache.Cache
	router *gin.Engine
	port   string
}

func New(port string, db *gorm.DB, cache *cache.Cache) *Server {
	s := &Server{db: db, cache: cache, router: gin.New(), port: port}

	s.router.Use(cors.New(config.Cors()))

	s.router.Use(sentrygin.New(sentrygin.Options{
		Repanic: true, // чтобы следующий recovery тоже сработал
	}))
	s.router.Use(gin.Recovery()) // panic → 500
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
