package routes

import (
	"net/http"

	"github.com/MochammadQemalFirza/echo_web/controllers"
	"github.com/labstack/echo"
)

func InitRouter() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Pegawai API !")
	})

	e.GET("/pegawai", controllers.FetchPegawai)
	e.GET("/pegawai/:id", controllers.FetchPegawaiID)
	e.POST("/pegawai", controllers.StorePegawai)
	e.PUT("/pegawai/:id", controllers.UpdatePegawai)
	e.DELETE("/pegawai", controllers.DeletePegawai)
	e.GET("/generate-hash/:passwrod", controllers.GenerateHashPassword)
	return e

}
