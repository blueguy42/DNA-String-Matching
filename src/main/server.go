package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func index(c echo.Context) error {
	return c.String(http.StatusOK, "halaman home")
	//ini akan jadi halaman home
}

func predict(c echo.Context) error {
	//semua logic predict ada disini, panggil API untuk predict di sini
	return c.String(http.StatusOK, "API predict")
}

func history(c echo.Context) error {
	//ini akan jadi halaman history
	return c.String(http.StatusOK, "halaman history")
}

func indexDisease(c echo.Context) error {
	//ini akan jadi API penambahan disease
	return c.String(http.StatusOK, "halaman add")
}

func addDisease(c echo.Context) error {
	//ini akan jadi API penambahan disease
	name := c.FormValue("disease")
	file, err := c.FormFile("dnasequence")

	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}

	fmt.Print(name, src)

	// dna, err = ioutil.ReadFile(src)
	// if err != nil {
	// 	return err
	// }
	return c.String(http.StatusOK, "API add")
}

func main() {
	fmt.Println("Hello, World !")
	e := echo.New()

	e.GET("/", index)
	e.POST("/", predict)
	e.GET("/history", history)
	e.GET("/add", indexDisease)
	e.POST("/add", addDisease)
	e.Start(":8000")
}
