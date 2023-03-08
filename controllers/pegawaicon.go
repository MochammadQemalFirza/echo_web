package controllers

import (
	"net/http"

	"github.com/MochammadQemalFirza/echo_web/models"
	"github.com/labstack/echo"
)

type Pegawai struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Position string `json:"position"`
}

func FetchPegawai(c echo.Context) error {
	result, err := models.FetchPegawai()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func StorePegawai(c echo.Context) error {
	// name := c.FormValue("name")
	// age := c.FormValue("age")
	// ages, _ := strconv.Atoi(age)
	// position := c.FormValue("position")
	// result, err := models.StorePegawai(name, ages, position)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, err.Error())
	// }
	// return c.JSON(http.StatusOK, result)
	var req Pegawai
	if err := c.Bind(&req); err != nil {
		return err
	}
	result, err := models.StorePegawai(req.Name, req.Age, req.Position)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, result)
}
