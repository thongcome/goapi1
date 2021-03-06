package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type inventory struct {
	id     string
	status string
	name   string
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

func getTodosHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[int]*inventory{
		1: &inventory{
			id:     "001",
			status: "processing",
			name:   "Notebook",
		},
	})

}

func helloHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "hello you ",
	})
}
