package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

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

	e.GET("/getTodos/:id", getByIDHandler)
	e.POST("/getTodos", createTodosHandler)
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
func createTodosHandler(c echo.Context) error {
	items := []*inventory{}
	// if err := e.JSON(http.StatusBadRequest,map[string]string {"error":err.Error()})
	for i := 1; i < 6; i++ {
		str := "00"
		str2 := strconv.Itoa(i)
		str = str + str2
		item := &inventory{
			ID:     str,
			Status: "processing",
			Name:   "Notebook",
		}
		items = append(items, item)
	}

	return c.JSON(http.StatusCreated, items)
}

func getTodosHandler(c echo.Context) error {
	items := []*inventory{}

	for i := 1; i < 6; i++ {
		str := "00"
		str2 := strconv.Itoa(i)
		str = str + str2
		item := &inventory{
			ID:     str,
			Status: "processing",
			Name:   "Notebook",
		}
		items = append(items, item)
	}
	// items = append(items, item)

	return c.JSON(http.StatusOK, items)

}

func getByIDHandler(c echo.Context) error {
	var id int
	err := echo.PathParamsBinder(c).Int("id", &id).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	items := []*inventory{}

	for i := 1; i < 6; i++ {
		str := "00"
		str2 := strconv.Itoa(i)
		str = str + str2
		item := &inventory{
			ID:     str,
			Status: "processing" + str2,
			Name:   "Notebook" + str2,
		}
		items = append(items, item)
	}
	t := items[id]
	// if !ok {
	// 	return c.JSON(http.StatusOK, map[int]*inventory{})
	// }

	// items = append(items, item)

	return c.JSON(http.StatusOK, t)

}

func helloHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "hello you ",
	})
}
