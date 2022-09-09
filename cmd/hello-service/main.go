package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID int `json:"id"`
}

var (
	users      = []*User{}
	default_id = 1
)

func createUser(c echo.Context) error {
	u := &User{
		ID: default_id,
	}

	users = append(users, u)

	err := c.Bind(u)
	if err != nil {
		return err
	}

	return nil

}
func main() {

	// helloHandler := func(w http.ResponseWriter, req *http.Request) {

	// 	io.WriteString(w, "Hello world \n")
	// }

	// http.HandleFunc("/hello", helloHandler)

	e := echo.New()

	e.POST("/users/create", createUser)

	e.GET("/users/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct {
			Length int `json:"length"`
		}{Length: len(users)})
	})

	log.Fatal(e.Start(":8000"))

}
