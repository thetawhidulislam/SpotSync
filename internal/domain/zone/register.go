package zone

import (
	"spotsync/internal/auth"
	"spotsync/internal/config"
	"spotsync/internal/middlewares"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB, cfg *config.Config) {
	repo := NewRepository(db)
	svc := NewService(repo)
	zoneHandler := NewHandler(svc)

	jwtService := auth.NewJWTService(cfg.JwtSecret)

	api := e.Group("/api/v1")

	api.POST(
		"/zones",
		zoneHandler.CreateZone,
		middlewares.AuthMiddleware(jwtService),
		middlewares.AdminMiddleware(),
	)
}
