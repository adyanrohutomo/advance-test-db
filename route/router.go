package route

import (
	"test-database/api"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/detail-order", api.DetailOrder)

	return e
}
