package zone

import (
	"errors"
	"github.com/labstack/echo/v5"
	"net/http"
	"spotsync/internal/domain/zone/dto"
	"spotsync/internal/httpresponse"
	"strconv"
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
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Validation failed",
			Details: err.Error(),
		})
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
func (h *handler) GetZone(c *echo.Context) error {
	zone, err := h.service.GetZone()
	if err != nil {
		return zoneErrorResponse(c, err)
	}
	return c.JSON(http.StatusOK, dto.ListAPIResponse{
		Success: true,
		Message: "Parking zones retrieved successfully",
		Data:    zone,
	})
}

func (h *handler) GetZoneByID(c *echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid Zone id",
			Details: err.Error(),
		})
	}

	response, err := h.service.GetZoneByID(uint(id))
	if err != nil {
		return zoneErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Parking zones retrieved successfully",
		Data:    *response,
	})
}
func (h *handler) UpdateZone(c *echo.Context) error {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid zone id")
	}

	var req dto.UpdateZoneRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	response, err := h.service.UpdateZone(uint(id), req)
	if err != nil {
		return zoneErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Parking zone updated successfully",
		Data:    *response,
	})
}
func (h *handler) DeleteZone(c *echo.Context) error {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid zone id")
	}

	err = h.service.DeleteZone(uint(id))
	if err != nil {
		return zoneErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Parking zone deleted successfully",
	})
}
