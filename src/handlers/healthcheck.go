package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"

	"go-api/src/models"
	"go-api/src/services"
)

// Handler defines the interface for healthcheck API handlers
//
// This interface is used to define all available healthcheck API methods, such as GetAPIStatus.
// Each method should be associated with an HTTP route in the implementation.
// The interface itself does not directly contribute to the Swagger documentation but
// serves as the blueprint for the handler implementation.
type HealthcheckHandler interface {
	// GetAPIStatus fetches the status of the API.
	GetAPIStatus(e echo.Context) error
}

// Params defines the dependencies that the healthcheck module needs
type Params struct {
	fx.In

	HealthcheckService services.HealthcheckService
}

type healthcheckHandler struct {
	hcService services.HealthcheckService
}

// New injects the healthcheck service
// into handler
func NewHealthcheckHandler(p Params) HealthcheckHandler {
	return &healthcheckHandler{
		hcService: p.HealthcheckService,
	}
}

// GetAPIStatus will return the status of the API
//
//	@Summary		Get API status
//	@Description	Get the status of the API
//	@Tags			healthcheck
//	@Produce		json
//	@Success		200	{object}	Status
//	@Failure		500	{string}	error
//	@Router			/ [get]
func (h *healthcheckHandler) GetAPIStatus(e echo.Context) error {
	onlineTime, err := h.hcService.OnlineSince()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	status := models.HealthcheckStatus{
		OnlineTime: onlineTime.String(),
	}

	return e.JSON(http.StatusOK, status)
}
