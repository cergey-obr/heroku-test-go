package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"os"
)

func main() {
	e := echo.New()

	e.POST("/", func(c echo.Context) error {
		request := c.Request()

		fmt.Println(request.Method)
		fmt.Println(request.Body)

		return c.String(http.StatusOK, "Hello, World!")
	})

	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("$PORT not set")
	}

	fmt.Println("hello")

	e.Logger.Fatal(e.Start(":" + port))
}
