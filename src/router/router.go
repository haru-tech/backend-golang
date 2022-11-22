package router

import (
	"cacao_tr_backend/controllers"

	"github.com/labstack/echo/v4"
)

func Init() {
	e := echo.New()

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	e.POST("/user", controllers.UserAuthentication)
	// e.PUT("/user/:id", controller.UpdateUser)
	// e.DELETE("/user/:id", controller.DeleteUser)

	e.Logger.Fatal(e.Start(":3000"))
}
