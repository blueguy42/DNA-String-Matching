package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

type (
	predictJSON struct {
		Name        string `json:"name"`
		Dna         string `json:"dna"`
		Diseasename string `json:"diseasename"`
		Mode        string `json:"mode"`
	}

	diseaseJSON struct {
		Names []string `json:"names"`
	}

	recordJSON struct {
		Date       string `json:"date"`
		Name       string `json:"name"`
		Disease    string `json:"disease"`
		Result     bool   `json:"result"`
		Similarity int    `json:"similarity"`
	}

	historyJSON struct {
		Records []recordJSON `json:"records"`
	}

	historyReqJSON struct {
		Date        string `json:"date"`
		Diseasename string `json:"diseasename"`
	}
)

func predict(c echo.Context) error {
	//API untuk predict di sini
	b := new(predictJSON)
	if err := c.Bind(b); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Print(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	defer db.Close()

	results, err := db.Query(fmt.Sprintf("SELECT dna FROM disease WHERE name='%s' LIMIT 1", b.Diseasename))
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	var diseasedna string
	for results.Next() {
		err = results.Scan(&diseasedna)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}
	var result bool
	if b.Mode == "kmp" {
		result = kmp(b.Dna, diseasedna)
	} else {
		result = boyermoore(b.Dna, diseasedna)
	}

	var similarity int

	similarity = lcsHighestSimilarity(b.Dna, diseasedna)

	if similarity >= 80 {
		result = true
	}

	var resultInt int
	if result {
		resultInt = 1
	} else {
		resultInt = 0
	}

	insert, err := db.Query(fmt.Sprintf("INSERT INTO history (date,name,disease,result,similarity) VALUES ('%s','%s','%s',%d,%d)", time.Now().Format("2006-01-02"), b.Name, b.Diseasename, resultInt, similarity))
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer insert.Close()

	r := &recordJSON{
		Date:       time.Now().Format("2006-01-02"),
		Name:       b.Name,
		Disease:    b.Diseasename,
		Result:     result,
		Similarity: similarity}

	return c.JSON(http.StatusOK, r)
}

func history(c echo.Context) error {
	//ini akan jadi halaman history
	b := new(historyReqJSON)
	if err := c.Bind(b); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Print(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	defer db.Close()
	var results *sql.Rows
	if b.Date == "" {
		results, err = db.Query(fmt.Sprintf("SELECT date,name,disease,result,similarity FROM history WHERE disease='%s'", b.Diseasename))
	} else if b.Diseasename == "" {
		results, err = db.Query(fmt.Sprintf("SELECT date,name,disease,result,similarity FROM history WHERE date='%s'", b.Date))
	} else {
		results, err = db.Query(fmt.Sprintf("SELECT date,name,disease,result,similarity FROM history WHERE date='%s' AND disease='%s'", b.Date, b.Diseasename))
	}

	if err != nil {
		fmt.Print(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	var record recordJSON
	var records []recordJSON
	for results.Next() {

		// for each row, scan the result into our tag composite object
		err = results.Scan(&record.Date, &record.Name, &record.Disease, &record.Result, &record.Similarity)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// and then print out the tag's Name attribute
		records = append(records, record)
	}

	history := &historyJSON{
		Records: records}
	return c.JSON(http.StatusOK, history)
}

func getDisease(c echo.Context) error {
	//ini akan jadi API untuk melihat disease yang ada

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Print(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	defer db.Close()

	results, err := db.Query("SELECT name FROM disease")

	if err != nil {
		fmt.Print(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	var name string
	var names []string
	for results.Next() {

		// for each row, scan the result into our tag composite object
		err = results.Scan(&name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// and then print out the tag's Name attribute
		names = append(names, name)
	}

	disease := &diseaseJSON{
		Names: names}
	return c.JSON(http.StatusOK, disease)
}

func addDisease(c echo.Context) error {
	//ini akan jadi API penambahan disease
	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Print(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	defer db.Close()

	insert, err := db.Query(fmt.Sprintf("INSERT INTO disease (name,dna) VALUES ('%s','%s')", jsonBody["name"], jsonBody["dna"]))

	if err != nil {
		fmt.Print(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	defer insert.Close()
	return c.String(http.StatusOK, "ok")
}

func home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.POST("/", home)
	e.POST("/predict", predict)
	e.POST("/history", history)
	e.GET("/get", getDisease)
	e.POST("/add", addDisease)
	e.Start(":" + os.Getenv("PORT"))
}
