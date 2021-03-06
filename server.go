package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type inventory struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Name   string `json:"name"`
}

func main() {
	e := echo.New()
	e.GET("/hello", helloHandler)

	e.GET("/getTodos", getTodosHandler)
	// e.Start(":1234")
	port := os.Getenv("PORT")
	log.Println("port", port)

	e.Start(":" + port)

}

// func getTodo2Handler(c echo.Context) error {
// 	item := []*inventory{}

// 	for
// 	return nil
// }

func getTodosHandler(c echo.Context) error {
	items := []*inventory{
		&inventory{
			ID:     "001",
			Status: "processing",
			Name:   "Notebook",
		},
		&inventory{
			ID:     "001",
			Status: "processing",
			Name:   "Notebook",
		},
		&inventory{
			ID:     "001",
			Status: "processing",
			Name:   "Notebook",
		},
	}

	return c.JSON(http.StatusOK, items)

}

func helloHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "hello you ",
	})
}
