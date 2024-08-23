// Package api provides api handlers for the Lambda
package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handler the API Handler.
type Handler struct{}

// NewHandler creates a new API Handler.
func NewHandler() *Handler {
	return &Handler{}
}

// Hello handler for Hello request.
func (h *Handler) Hello(c echo.Context) error {
	c.Logger().Info("Hello!")
	return c.JSON(http.StatusOK, h.jsonResponseWithMessage("Hello!"))
}

// SayMyName request for the SayMyName handler.
type SayMyName struct {
	Name string `json:"name"`
}

// YourName handlers for your name request.
func (h *Handler) SayMyName(c echo.Context) error {
	r := c.Request()
	payload := SayMyName{}
	defer func() {
		err := r.Body.Close()
		if err != nil {
			c.Logger().Error("unable to close response body: %s", err)
		}
	}()
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, h.jsonResponseWithMessage("invalid request body provided"))
	}
	if payload.Name == "" {
		return c.JSON(http.StatusBadRequest, h.jsonResponseWithMessage("if you don't tell me I don't know your name"))
	}

	return c.JSON(http.StatusOK, h.jsonResponseWithMessage(fmt.Sprintf("Your name is %s", payload.Name)))
}

func (h *Handler) jsonResponseWithMessage(msg string) any {
	return struct {
		Message string `json:"message"`
	}{
		Message: msg,
	}
}
