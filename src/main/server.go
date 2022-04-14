package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func index(c echo.Context) error {
	return c.String(http.StatusOK, "hello from the website")
	//ini akan jadi halaman home
}

func predict(c echo.Context) error {
	//semua logic predict ada disini, panggil API untuk predict di sini
	return c.String(http.StatusOK, "hello from the website")
}

func history(c echo.Context) error {
	//ini akan jadi halaman history
	return c.String(http.StatusOK, "hello from the website")
}

func main() {
	fmt.Println("Hello, World !")
	e := echo.New()

	e.GET("/", index)
	e.POST("/", predict)
	e.GET("/history", history)

	e.Start(":8000")
}
