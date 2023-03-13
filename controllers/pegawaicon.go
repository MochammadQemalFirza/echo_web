package controllers

import (
	"net/http"
	"strconv"

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

func FetchPegawaiID(c echo.Context) error {
	id := c.Param("id")
	idconv, _ := strconv.Atoi(id)
	result, err := models.FetchPegawaiID(idconv)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func StorePegawai(c echo.Context) error {
	//form
	// name := c.FormValue("name")
	// age := c.FormValue("age")
	// ages, _ := strconv.Atoi(age)
	// position := c.FormValue("position")
	// result, err := models.StorePegawai(name, ages, position)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, err.Error())
	// }
	// return c.JSON(http.StatusOK, result)

	//json
	var req Pegawai
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	result, err := models.StorePegawai(req.Name, req.Age, req.Position)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, result)
}

func UpdatePegawai(c echo.Context) error {
	var req Pegawai
	id := c.Param("id")
	idconv, _ := strconv.Atoi(id)
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	result, err := models.UpdatePegawai(idconv, req.Name, req.Age, req.Position)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, result)
}

func DeletePegawai(c echo.Context) error {
	var req Pegawai
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	result, err := models.DeletePegawai(req.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())

	}
	return c.JSON(http.StatusOK, result)
}
