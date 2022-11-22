package controllers

import (
	"fmt"
	"net/http"

	"cacao_tr_backend/interfaces"

	"github.com/labstack/echo/v4"
)

func UserAuthentication(c echo.Context) error {
	// test

	ret := interfaces.IF01_01_OUT{
		Result:    "OK",
		SessionID: "1",
		Detail:    "none"}

	fmt.Println(ret)

	return c.JSON(http.StatusOK, ret)
}
