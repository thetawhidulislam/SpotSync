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

	api := e.Group("/api/v1/zones")

	api.POST(
		"/",
		zoneHandler.CreateZone,
		middlewares.AuthMiddleware(jwtService),
		middlewares.AdminMiddleware(),
	)
	api.GET("/", zoneHandler.GetZone)
	api.GET("/:id", zoneHandler.GetZoneByID)
	api.PATCH(
		"/:id",
		zoneHandler.UpdateZone,
		middlewares.AuthMiddleware(jwtService),
		middlewares.AdminMiddleware(),
	)

	api.DELETE(
		"/:id",
		zoneHandler.DeleteZone,
		middlewares.AuthMiddleware(jwtService),
		middlewares.AdminMiddleware(),
	)
}
