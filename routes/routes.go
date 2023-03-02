package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

func InitRouter(*echo.Echo) {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Test Echo Framework !")
	})

}
