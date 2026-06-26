package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	// "github.com/labstack/echo/v5/middleware"
	"gorm.io/gorm"
	"net/http"
	"spotsync/internal/config"
	// "spotsync/internal/domain/users"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func Start(db *gorm.DB, cfg *config.Config) {
	db.AutoMigrate()

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/health", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Server is running",
		})
	})
	e.Start(":" + cfg.PORT)
}
