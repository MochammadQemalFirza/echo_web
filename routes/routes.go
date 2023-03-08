package routes

import (
	"net/http"

	"github.com/MochammadQemalFirza/echo_web/controllers"
	"github.com/labstack/echo"
)

func InitRouter() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Test Echo Framework !")
	})

	e.GET("/pegawai", controllers.FetchPegawai)
	e.POST("/pegawai", controllers.StorePegawai)

	return e

}
