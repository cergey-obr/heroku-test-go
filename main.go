package main

import (
	"fmt"
	"github.com/labstack/echo"
	"os"
)

func main() {
	e := echo.New()

	e.POST("/", func(c echo.Context) error {
		//payload := map[string]interface{}{}
		data := c.FormValue("payload")
		fmt.Println(data)

		return c.JSON(200, echo.Map{})
	})

	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("$PORT not set")
	}

	fmt.Println("hello")

	e.Logger.Fatal(e.Start(":" + port))
}
