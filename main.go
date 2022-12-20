package main

import (
	"net/http"
	"github.com/labstack/echo"
)

func main(){
	e := echo.New()

	e.GET("/hello", Hello())
	e.GET("/api/hello", ApiHelloGet())

	e.Start(":8080")
}

func Hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "4433: 12/20 exp10 ver2")
	}
}

func ApiHelloGet() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"studentId": "4433", "message": "12/20 exp10 ver2"})
	}
}