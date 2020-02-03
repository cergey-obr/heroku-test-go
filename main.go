package main

import (
	"fmt"
	"github.com/labstack/echo"
	"os"
)

func main() {
	e := echo.New()

	e.POST("/", func(c echo.Context) error {
		m := echo.Map{}
		if err := c.Bind(&m); err != nil {
			return err
		}

		fmt.Println(m)

		return c.JSON(200, m)
	})

	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("$PORT not set")
	}

	fmt.Println("hello")

	e.Logger.Fatal(e.Start(":" + port))
}
