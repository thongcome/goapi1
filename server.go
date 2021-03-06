package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
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
	e.GET("/all", getAllDBHandler)

	e.GET("/getTodos/:id", getByIDHandler)
	e.GET("/getDBByID/:id", getDBByIDHandler)
	e.POST("/getTodos", createTodosHandler)
	e.POST("/insertData", insertDataHandler)
	e.POST("/updateData", updateDataHandler)
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

func getDBByIDHandler(c echo.Context) error {
	var id int
	err := echo.PathParamsBinder(c).Int("id", &id).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT id, title, status FROM todos where id=$1")
	if err != nil {
		log.Fatal("can'tprepare query one row statment", err)
	}

	rowId := id
	row := stmt.QueryRow(rowId)

	var title, status string

	err = row.Scan(&id, &title, &status)
	if err != nil {
		log.Fatal("can't Scan row into variables", err)
	}
	item := &inventory{
		ID:     strconv.Itoa(id),
		Status: title,
		Name:   status,
	}

	fmt.Println("one row", id, title, status)
	return c.JSON(http.StatusOK, item)

}

func getAllDBHandler(c echo.Context) error {

	items := []*inventory{}
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT id, title, status FROM todos order by id desc")
	if err != nil {
		log.Fatal("can't prepare query all todos statment", err)
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal("can't query all todos", err)
	}
	for rows.Next() {
		var id int
		var title, status string
		err := rows.Scan(&id, &title, &status)
		if err != nil {
			log.Fatal("can't Scan row into variable", err)
		}
		fmt.Println(id, title, status)
		item := &inventory{
			ID:     strconv.Itoa(id),
			Status: title,
			Name:   status,
		}

		items = append(items, item)
	}

	fmt.Println("query all todos success")
	return c.JSON(http.StatusOK, items)

}

func insertDataHandler(c echo.Context) error {
	item := inventory{}
	err := c.Bind(&item)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	row := db.QueryRow("INSERT INTO todos (title, status) values ($1, $2)  RETURNING id", item.Name, item.Status)
	var id int
	err = row.Scan(&id)
	if err != nil {
		fmt.Println("can't scan id", err)
		return err
	}

	fmt.Println("insert todo success id : ", id)
	return c.JSON(http.StatusOK, id)
}

func updateDataHandler(c echo.Context) error {

	item := inventory{}
	err := c.Bind(&item)
	if item.ID == "" {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	// if err != nil {
	// 	log.Fatal("Connect to database error", err)
	// }
	// defer db.Close()

	// stmt, err := db.Prepare("UPDATE todos SET status=$2 , name=$3 WHERE id=$1;")

	// if err != nil {
	// 	log.Fatal("can't prepare statment update", err)
	// }
	// int1, err := strconv.ParseInt(item.ID, 6, 12)

	// if _, err := stmt.Exec(int1, item.Status, item.Name); err != nil {
	// 	log.Fatal("error execute update ", err)
	// }

	// fmt.Println("insert todo success id : ", int1)
	return c.JSON(http.StatusOK, item)

}

func helloHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "hello you ",
	})
}
