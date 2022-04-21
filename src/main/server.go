package main

import (
	"encoding/json"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

type (
	predictJSON struct {
		name    string `json:"name"`
		dna     string `json:"dna"`
		disease string `json:"disease"`
	}

	addJSON struct {
		name string `json:"name"`
		dna  string `json:"dna"`
	}
)

const HOST = "127.0.0.1"
const PORT = "3306"
const USER = "root"
const PASS = "pass"
const DBNAME = "dnamatching"

func predict(c echo.Context) error {
	//semua logic predict ada disini, panggil API untuk predict di sini
	return c.String(http.StatusOK, "API predict")
}

func history(c echo.Context) error {
	//ini akan jadi halaman history
	return c.String(http.StatusOK, "halaman history")
}

func getDisease(c echo.Context) error {
	//ini akan jadi API untuk melihat disease yang ada
	return c.String(http.StatusOK, "halaman add")
}

func addDisease(c echo.Context) error {
	//ini akan jadi API penambahan disease
	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		return c.String(http.StatusBadRequest, "failed")
	}
	s := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", USER, PASS, HOST, PORT, DBNAME)
	db, err := sql.Open("mysql", fmt.Sprintf("%s@tcp(%s:%s)/%s", USER, HOST, PORT, DBNAME))
	fmt.Print(s)
	if err != nil {
		fmt.Print(err)
		return c.String(http.StatusInternalServerError, "connection to database failed")
	}

	defer db.Close()

	
	insert, err := db.Query(fmt.Sprintf("INSERT INTO disease (name,dna) VALUES ( \"%s\", \"%s\" )", jsonBody["name"], jsonBody["dna"]))

	if err != nil {
		fmt.Print(err)
		return c.String(http.StatusInternalServerError, "failed to insert")
	}

	defer insert.Close()
	return c.String(http.StatusOK, "ok")
}

func main() {
	fmt.Println("Hello, World !")
	e := echo.New()

	e.POST("/predict", predict)
	e.GET("/history", history)
	e.GET("/get", getDisease)
	e.POST("/add", addDisease)
	e.Start(":8000")
}
