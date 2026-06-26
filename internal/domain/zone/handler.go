package zone

import (
	"errors"
	"net/http"
	"spotsync/internal/domain/zone/dto"
	"spotsync/internal/httpresponse"

	"github.com/labstack/echo/v5"
)

type handler struct {
	service *service
}

func NewHandler(s *service) *handler {
	return &handler{service: s}
}
func zoneErrorResponse(c *echo.Context, err error) error {
	if errors.Is(err, ErrZoneNotFound) {
		return c.JSON(http.StatusNotFound, httpresponse.Error{
			Code:    http.StatusNotFound,
			Message: "Zone not found",
		})
	}
	return c.JSON(http.StatusInternalServerError, httpresponse.Error{
		Code:    http.StatusInternalServerError,
		Message: "Something went wrong",
		Details: err.Error(),
	})
}

func (h *handler) CreateZone(c *echo.Context) error {

	var req dto.CreateZoneRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	response, err := h.service.CreateZone(req)
	if err != nil {
		return zoneErrorResponse(c, err)
	}

	return c.JSON(http.StatusCreated, dto.APIResponse{
		Success: true,
		Message: "Parking zone created successfully",
		Data:    *response,
	})
}