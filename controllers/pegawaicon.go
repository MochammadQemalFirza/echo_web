package controllers

import (
	"net/http"

	"github.com/MochammadQemalFirza/echo_web/models"
	"github.com/labstack/echo"
)

func FetchPegawai(c echo.Context) error {
	result, err := models.FetchPegawai()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}
