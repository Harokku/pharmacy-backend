package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

//returnError helper to return string containing error and log it
//
//c: echo.Context.
//
//s int: http.status.
//
//e error: error to return.
//
//m string: error string.
func returnError(c echo.Context, s int, e error, m string) error {
	c.Logger().Errorf("%v: %v", m, e)
	return c.String(s, fmt.Sprintf("%v: %v", m, e))
}
