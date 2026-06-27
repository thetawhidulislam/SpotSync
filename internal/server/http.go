package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"gorm.io/gorm"
	"net/http"
	"spotsync/internal/config"
	"spotsync/internal/domain/reservation"
	"spotsync/internal/domain/user"
	"spotsync/internal/domain/zone"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func Start(db *gorm.DB, cfg *config.Config) {
	db.AutoMigrate(&user.User{}, &zone.Zone{}, &reservation.Reservation{})

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Server is running",
		})
	})

	e.Use(middleware.RequestLogger())
	user.RegisterRoutes(e, db, cfg)
	zone.RegisterRoutes(e, db, cfg)
	reservation.RegisterRoutes(e, db, cfg)
	e.Start(":" + cfg.PORT)

}
