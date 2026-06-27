package reservation

import (
	"spotsync/internal/auth"
	"spotsync/internal/config"
	"spotsync/internal/domain/zone"
	"spotsync/internal/middlewares"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB, cfg *config.Config) {
	orderRepo := NewRepository(db)
	zoneRepo := zone.NewRepository(db)

	svc := NewService(orderRepo, zoneRepo)
	handler := NewHandler(svc)

	jwtService := auth.NewJWTService(cfg.JwtSecret)

	api := e.Group("/api/v1/reservations", middlewares.AuthMiddleware(jwtService))

	api.POST("", handler.CreateReservation)
	api.GET("/my-reservations", handler.MyReservations)
}
