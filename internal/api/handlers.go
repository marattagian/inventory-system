package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/marattagian/inventory-system/internal/api/dtos"
	"github.com/marattagian/inventory-system/internal/service"
)

type responseMessage struct {
	Message string `json:"message"`
}

func (a *API) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.RegisterUser{}

	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "invalid request"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	err = a.serv.RegisterUser(ctx, params.Email, params.Name, params.Password)
	if err != nil {
		if err == service.ErrUserAlreadyExists {
			return c.JSON(http.StatusConflict, responseMessage{Message: err.Error()})
		}

		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "internal server error"})
	}

	return c.JSON(http.StatusCreated, nil)
}
