package controllers

import (
	"net/http"

	"github.com/MochammadQemalFirza/echo_web/helper"
	"github.com/labstack/echo"
)

func GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")
	hash, _ := helper.HashPassword(password)
	return c.JSON(http.StatusOK, hash)
}
