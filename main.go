package main

import (
	"fmt"
	"github.com/labstack/echo"
	"io/ioutil"
	"os"
)

func main() {
	e := echo.New()

	e.POST("/", func(c echo.Context) error {
		body := c.Request().Body
		b, _ := ioutil.ReadAll(body)
		fmt.Println(b)

		return c.JSON(200, echo.Map{})
	})

	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("$PORT not set")
	}

	fmt.Println("hello")

	e.Logger.Fatal(e.Start(":" + port))
}
