package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
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
)

const HOST = "127.0.0.1"
const PORT = "3306"
const USER = "root"
const PASS = "pass"
const DBNAME = "dnamatching"

func predict(c echo.Context) error {
	//API untuk predict di sini
	b := new(predictJSON)
	if err := c.Bind(b); err != nil {
		return c.String(http.StatusBadRequest, "failed")
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s@tcp(%s:%s)/%s", USER, HOST, PORT, DBNAME))

	if err != nil {
		fmt.Print(err)
		return c.String(http.StatusInternalServerError, "connection to database failed")
	}

	defer db.Close()

	results, err := db.Query(fmt.Sprintf("SELECT dna FROM disease WHERE name=%s LIMIT 1", b.Diseasename))
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

	similarity := lcsHighestSimilarity(b.Dna, diseasedna)

	insert, err := db.Query(fmt.Sprintf("INSERT INTO history (date,name,disease,result,similarity) VALUES (%s,%s,%s,%v,%d)", time.Now().Format("01-02-2006"), b.Name, b.Diseasename, result, similarity))
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer insert.Close()

	return c.String(http.StatusOK, "ok")
}

func history(c echo.Context) error {
	//ini akan jadi halaman history
	db, err := sql.Open("mysql", fmt.Sprintf("%s@tcp(%s:%s)/%s", USER, HOST, PORT, DBNAME))

	if err != nil {
		fmt.Print(err)
		return c.String(http.StatusInternalServerError, "connection to database failed")
	}

	defer db.Close()

	results, err := db.Query("SELECT date,name,disease,result FROM history")

	if err != nil {
		fmt.Print(err)
		return c.String(http.StatusInternalServerError, "failed to query")
	}

	var record recordJSON
	var records []recordJSON
	for results.Next() {

		// for each row, scan the result into our tag composite object
		err = results.Scan(&record.Date, &record.Name, &record.Disease, &record.Result)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// and then print out the tag's Name attribute
		records = append(records, record)
		fmt.Println(record)
	}

	history := &historyJSON{
		Records: records}
	return c.JSON(http.StatusOK, history)
}

func getDisease(c echo.Context) error {
	//ini akan jadi API untuk melihat disease yang ada

	db, err := sql.Open("mysql", fmt.Sprintf("%s@tcp(%s:%s)/%s", USER, HOST, PORT, DBNAME))

	if err != nil {
		fmt.Print(err)
		return c.String(http.StatusInternalServerError, "connection to database failed")
	}

	defer db.Close()

	results, err := db.Query("SELECT name FROM disease")

	if err != nil {
		fmt.Print(err)
		return c.String(http.StatusInternalServerError, "failed to query")
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
		fmt.Println(name)
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
		return c.String(http.StatusBadRequest, "failed")
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s@tcp(%s:%s)/%s", USER, HOST, PORT, DBNAME))

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
