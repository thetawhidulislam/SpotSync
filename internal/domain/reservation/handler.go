package reservation

import (
	"errors"
	"github.com/labstack/echo/v5"
	"net/http"
	"spotsync/internal/domain/reservation/dto"
	"spotsync/internal/domain/zone"
	"spotsync/internal/httpresponse"
	"strconv"
)

type handler struct {
	service *service
}

func NewHandler(s *service) *handler {
	return &handler{service: s}
}
func getCurrentUserID(c *echo.Context) (uint, bool) {
	userId, ok := c.Get("user_id").(uint)
	return userId, ok
}
func orderErrorResponse(c *echo.Context, err error) error {
	if errors.Is(err, ErrOrderNotFound) {
		return c.JSON(http.StatusNotFound, httpresponse.Error{
			Code:    http.StatusNotFound,
			Message: "Order not found",
		})
	}
	if errors.Is(err, zone.ErrZoneNotFound) {
		return c.JSON(http.StatusNotFound, httpresponse.Error{
			Code:    http.StatusNotFound,
			Message: "Zone not found",
		})
	}
	if errors.Is(err, ErrZoneFull) {
		return c.JSON(http.StatusConflict, httpresponse.Error{
			Code:    http.StatusConflict,
			Message: "Parking zone is full",
		})
	}
	if errors.Is(err, ErrOrderAlreadyCancelled) {
		return c.JSON(http.StatusConflict, httpresponse.Error{
			Code:    http.StatusConflict,
			Message: "Order is already cancelled",
		})
	}
	if errors.Is(err, ErrForbiddenOrderAccess) {
		return c.JSON(http.StatusForbidden, httpresponse.Error{
			Code:    http.StatusForbidden,
			Message: "You do not own this order",
		})
	}
	return c.JSON(http.StatusInternalServerError, httpresponse.Error{
		Code:    http.StatusInternalServerError,
		Message: "Something went wrong",
		Details: err.Error(),
	})
}

func (h *handler) CreateReservation(c *echo.Context) error {
	userId, ok := getCurrentUserID(c)
	if !ok {
		return c.JSON(http.StatusUnauthorized, httpresponse.Error{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		})
	}

	var req dto.CreateReservationRequest
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

	response, err := h.service.CreateOrder(userId, req)
	if err != nil {
		return orderErrorResponse(c, err)
	}

	return c.JSON(http.StatusCreated, dto.APIResponse{
		Success: true,
		Message: "Reservation confirmed successfully",
		Data:    *response,
	})
}

func (h *handler) MyReservations(c *echo.Context) error {
	userId, ok := getCurrentUserID(c)
	if !ok {
		return c.JSON(http.StatusUnauthorized, httpresponse.Error{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		})
	}

	response, err := h.service.GetMyReservations(userId)
	if err != nil {
		return orderErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, dto.ListAPIResponse{
		Success: true,
		Message: "My reservations retrieved successfully",
		Data:    response,
	})
}
func (h *handler) CancelReservation(c *echo.Context) error {

	userId, ok := getCurrentUserID(c)
	if !ok {
		return c.JSON(http.StatusUnauthorized, httpresponse.Error{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		})
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid reservation id",
		})
	}

	err = h.service.CancelReservation(userId, uint(id))
	if err != nil {
		return orderErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Reservation cancelled successfully",
	})
}
